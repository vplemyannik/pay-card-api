package tracer

import (
	"context"
	"fmt"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
	"io"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"

	"github.com/ozonmp/pay-card-api/internal/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer - returns new tracer.
func NewTracer(cfg *config.Config) (io.Closer, error) {
	ctx := context.Background()
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}
	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("failed init jaeger: %v", err))

		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	logger.InfoKV(ctx, "Traces started")

	return closer, nil
}
