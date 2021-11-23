package middleware

import (
	"context"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func parseLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.InfoLevel, true
	case "warn":
		return zapcore.WarnLevel, true
	case "error":
		return zapcore.ErrorLevel, true
	default:
		return zapcore.DebugLevel, false
	}
}

func LogLevelSwitchInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			levels := md.Get("log-level")

			logger.InfoKV(ctx, "got log level", "levels", levels)

			if len(levels) > 0 {
				if parsedLevel, ok := parseLevel(levels[0]); ok {
					newLogger := logger.CloneWithLevel(ctx, parsedLevel)
					ctx = logger.AttachLogger(ctx, newLogger)
				}
			}
		}

		return handler(ctx, req)
	}
}
