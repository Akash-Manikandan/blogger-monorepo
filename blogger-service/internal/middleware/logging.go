package middleware

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Global Zap logger
var logger *zap.Logger

func init() {
	cfg := zap.NewDevelopmentConfig()

	// Custom log format for performance & readability
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Colorized output
	cfg.EncoderConfig.EncodeTime = magentaTimeEncoder                // Custom magenta timestamp
	cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic("Failed to initialize Zap logger")
	}
	defer logger.Sync()
}

// ANSI escape code for magenta color
const magenta = "\033[35m"
const reset = "\033[0m"

// Custom time encoder to color timestamp in magenta
func magentaTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s%s%s", magenta, t.Format("2006-01-02T15:04:05.000Z0700"), reset))
}

// Extract headers from gRPC metadata
func extractHeaders(ctx context.Context) map[string]string {
	headers := make(map[string]string)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, values := range md {
			headers[key] = values[0] // Log only the first value
		}
	}
	return headers
}

// Unary Interceptor (for non-streaming gRPC requests)
func UnaryLoggingInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	startTime := time.Now()
	headers := extractHeaders(ctx)

	logger.Info("📩 gRPC Request",
		zap.String("Method", info.FullMethod),
		zap.Time("Timestamp", startTime),
		zap.Any("Headers", headers),
		zap.Any("Request", req), // Direct logging without JSON formatting
	)

	// Call gRPC handler
	resp, err := handler(ctx, req)

	duration := time.Since(startTime)
	logger.Info("📤 gRPC Response",
		zap.String("Method", info.FullMethod),
		zap.Duration("Duration", duration),
		zap.Any("Response", resp),
	)

	if err != nil {
		logger.Error("🚨 gRPC Error",
			zap.String("Method", info.FullMethod),
			zap.Error(err),
		)
	}

	return resp, err
}

// Stream Interceptor (for streaming gRPC calls)
func StreamLoggingInterceptor(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	startTime := time.Now()

	logger.Info("🔄 gRPC Streaming Started",
		zap.String("Method", info.FullMethod),
		zap.Time("Timestamp", startTime),
	)

	err := handler(srv, ss)

	if err != nil {
		logger.Error("🚨 gRPC Stream Error",
			zap.String("Method", info.FullMethod),
			zap.Error(err),
		)
	} else {
		logger.Info("✅ gRPC Stream Completed",
			zap.String("Method", info.FullMethod),
			zap.Duration("Duration", time.Since(startTime)),
		)
	}

	return err
}
