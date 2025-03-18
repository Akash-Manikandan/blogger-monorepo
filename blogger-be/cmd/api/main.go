package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Akash-Manikandan/blogger-be/internal/config"
	"github.com/Akash-Manikandan/blogger-be/internal/middleware"
	"github.com/Akash-Manikandan/blogger-be/internal/routes/blogs"
	"github.com/Akash-Manikandan/blogger-be/internal/routes/health"
	"github.com/Akash-Manikandan/blogger-be/internal/routes/users"
	"github.com/Akash-Manikandan/blogger-be/internal/utils"

	blogPb "github.com/Akash-Manikandan/blogger-be/proto/blog/v1"
	healthPb "github.com/Akash-Manikandan/blogger-be/proto/health/v1"
	userPb "github.com/Akash-Manikandan/blogger-be/proto/user/v1"

	"google.golang.org/grpc"
)

func main() {
	config.LoadEnv()
	// Load configuration
	cfg := config.Load()

	utils.LogAllMethods()

	// Initialize database connection
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database schemas
	// if err := models.AutoMigrate(db); err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }

	// Initialize gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.JWTAuthInterceptor,        // JWT Auth middleware
			middleware.UnaryValidatorInterceptor, // Validation middleware
			middleware.UnaryLoggingInterceptor,   // Logging middleware
		),
		grpc.StreamInterceptor(middleware.StreamLoggingInterceptor),
	)
	healthServer := health.NewBlogHealthServer(db)
	usersServer := users.UserServerRegister(db)
	blogsServer := blogs.NewBlogServer(db)
	healthPb.RegisterHealthServiceServer(grpcServer, healthServer)
	userPb.RegisterUserServiceServer(grpcServer, usersServer)
	blogPb.RegisterBlogServiceServer(grpcServer, blogsServer)

	log.Printf("Starting gRPC server on port %s", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
