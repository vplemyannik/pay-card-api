package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/pay-card-api/internal/app/retranslator"
	"github.com/ozonmp/pay-card-api/internal/app/sender"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
	repo_cards "github.com/ozonmp/pay-card-api/internal/repo/cards"
	repo_cards_events "github.com/ozonmp/pay-card-api/internal/repo/cards_events"
	"github.com/ozonmp/pay-card-api/internal/server/middleware"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/ozonmp/pay-card-api/internal/api"
	"github.com/ozonmp/pay-card-api/internal/config"
	pb "github.com/ozonmp/pay-card-api/pkg/pay-card-api"
)

// GrpcServer is gRPC server
type GrpcServer struct {
	db        *sqlx.DB
	batchSize uint
}

// NewGrpcServer returns gRPC server with supporting of batch listing
func NewGrpcServer(db *sqlx.DB, batchSize uint) *GrpcServer {
	return &GrpcServer{
		db:        db,
		batchSize: batchSize,
	}
}

// Start method runs server
func (s *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gatewayAddr := fmt.Sprintf("%s:%v", cfg.Rest.Host, cfg.Rest.Port)
	grpcAddr := fmt.Sprintf("%s:%v", cfg.Grpc.Host, cfg.Grpc.Port)
	metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)

	gatewayServer := createGatewayServer(grpcAddr, gatewayAddr)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("Gateway server is running on %s", gatewayAddr))
		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running gateway server")
			cancel()
		}
	}()

	metricsServer := createMetricsServer(cfg)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("Metrics server is running on %s", metricsAddr))
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running metrics server")
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(cfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, fmt.Sprintf("Status server is running on %s", statusAdrr))
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running status server")
		}
	}()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_opentracing.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
			middleware.LogLevelSwitchInterceptor(),
		)),
	)

	repo := repo_cards.NewCardRepo(s.db)
	repoEvents := repo_cards_events.NewCardEventsRepo(s.db)
	producer, err := sender.NewSyncProducer([]string{"localhost:9094"})
	if err != nil {
		logger.ErrorKV(ctx, "Failed running kafka producer")
		return err
	}
	eventSender := sender.NewKafkaSender(producer)

	pb.RegisterPayCardApiServiceServer(grpcServer, api.NewTemplateAPI(repo, repoEvents))
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("Gateway server is running on %s", gatewayAddr))
		if err := grpcServer.Serve(l); err != nil {
			logger.ErrorKV(ctx, "Failed running gateway server")
		}
	}()

	retranslatorCfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  1,
		ConsumeSize:    2,
		ConsumeTimeout: 1 * time.Second,
		ProducerCount:  1,
		WorkerCount:    1,
		Repo:           repoEvents,
		Sender:         eventSender,
	}

	retranslator := retranslator.NewRetranslator(retranslatorCfg)
	go func() {
		logger.InfoKV(ctx, "Retranslator is running")
		retranslator.Start()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.InfoKV(ctx, "The service is ready to accept requests")
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		logger.InfoKV(ctx, fmt.Sprintf("ctx.Done: %v", done))
	}

	isReady.Store(false)

	if err := gatewayServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, "gatewayServer.Shutdown")
	} else {
		logger.InfoKV(ctx, "gatewayServer shut down correctly")
	}

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, "statusServer.Shutdown")
	} else {
		logger.InfoKV(ctx, "statusServer shut down correctly")
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, "metricsServer.Shutdown")
	} else {
		logger.InfoKV(ctx, "metricsServer shut down correctly")
	}

	retranslator.Close()
	logger.InfoKV(ctx, "retranslator shut down correctly")

	grpcServer.GracefulStop()
	logger.InfoKV(ctx, "grpcServer shut down correctly")

	return nil
}
