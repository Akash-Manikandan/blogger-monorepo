package middleware

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// UnaryValidatorInterceptor intercepts gRPC unary requests to validate fields
func UnaryValidatorInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	protoReq, ok := req.(proto.Message)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to cast request to proto.Message")
	}
	// Validate the request fields
	if err := protovalidate.Validate(protoReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	// Proceed with the actual handler
	return handler(ctx, req)
}
