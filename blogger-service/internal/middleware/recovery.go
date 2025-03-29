package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryPanicRecoveryInterceptor intercepts gRPC unary requests to recover from panics.
func UnaryPanicRecoveryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic.
			fmt.Printf("panic occurred: %v\n", r)
			// Return a standardized internal error.
			err = status.Errorf(codes.Internal, "internal server error")
		}
	}()
	// Continue with the request handler.
	return handler(ctx, req)
}

// StreamPanicRecoveryInterceptor intercepts gRPC stream requests to recover from panics.
func StreamPanicRecoveryInterceptor(
	srv any,
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic.
			fmt.Printf("panic occurred: %v\n", r)
			// Return a standardized internal error.
			err = status.Errorf(codes.Internal, "internal server error")
		}
	}()
	// Continue with the stream handler.
	return handler(srv, stream)
}
