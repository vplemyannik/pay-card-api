package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/ozonmp/pay-card-api/internal/repo"

	pb "github.com/ozonmp/pay-card-api/pkg/pay-card-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "omp_template_api_template_not_found_total",
		Help: "Total number of templates that were not found",
	})
)

type cardAPI struct {
	pb.UnimplementedPayCardApiServiceServer
	repo repo.Repo
}

func NewTemplateAPI(r repo.Repo) pb.PayCardApiServiceServer {
	return &cardAPI{repo: r}
}

func (a cardAPI) CreateCard(ctx context.Context, req *pb.CreateCardV1Request) (*pb.CreateCardV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateCardV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("CreateCard request happens")

	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (a cardAPI) RemoveCard(ctx context.Context, req *pb.RemoveCardV1Request) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveCardV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("RemoveCard request happens")

	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (a cardAPI) DescribeCard(ctx context.Context, req *pb.DescribeCardV1Request) (*pb.Card, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeCardV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("DescribeCard request happens")

	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (a cardAPI) ListCard(ctx context.Context, empty *emptypb.Empty) (*pb.ListCardV1Response, error) {

	log.Debug().Msg("ListCard request happens")

	return nil, status.Error(codes.Unimplemented, "Not implemented")
}
