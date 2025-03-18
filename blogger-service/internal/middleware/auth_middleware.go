package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	authpb "github.com/Akash-Manikandan/blogger-service/proto/auth/v1" // Import the generated proto extension
)

type contextKey string

const UserIDKey contextKey = "userId"

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func extractJWT(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata not provided")
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token not provided")
	}

	tokenParts := strings.Split(authHeader[0], " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", status.Errorf(codes.Unauthenticated, "invalid authorization format")
	}

	return tokenParts[1], nil
}

func validateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "invalid token claims")
	}

	// Check token expiration
	if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
		return "", status.Errorf(codes.Unauthenticated, "token expired")
	}

	// Extract the user ID from "sub" claim
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "user ID not found in token")
	}

	return userID, nil
}

func methodRequiresAuth(fullMethod string) (bool, error) {
	fullMethod = strings.TrimPrefix(fullMethod, "/")

	parts := strings.Split(fullMethod, "/")
	if len(parts) != 2 {
		return false, nil
	}

	service, method := parts[0], parts[1]
	formattedMethod := fmt.Sprintf("%s.%s", service, method)

	descriptor, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(formattedMethod))
	if err != nil {
		return false, nil
	}

	methodDescriptor, ok := descriptor.(protoreflect.MethodDescriptor)
	if !ok {
		return false, status.Errorf(codes.Internal, "invalid method descriptor")
	}

	ext := proto.GetExtension(methodDescriptor.Options(), authpb.E_RequiresAuth)
	if ext == nil {
		return false, nil
	}

	requiresAuth, ok := ext.(bool)
	if !ok {
		return false, status.Errorf(codes.Internal, "invalid auth extension type")
	}

	return requiresAuth, nil
}

// JWT Auth Interceptor
func JWTAuthInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	authRequired, err := methodRequiresAuth(info.FullMethod)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check auth requirement: %v", err)
	}

	if authRequired {
		tokenString, err := extractJWT(ctx)
		if err != nil {
			return nil, err
		}

		userID, err := validateJWT(tokenString)
		if err != nil {
			return nil, err
		}

		// Store userID in context for service logic
		ctx = context.WithValue(ctx, UserIDKey, userID)
	}

	return handler(ctx, req)
}
