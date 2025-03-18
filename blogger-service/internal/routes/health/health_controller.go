package health

import (
	"context"

	pb "github.com/Akash-Manikandan/blogger-service/proto/health/v1"

	"gorm.io/gorm"
)

type BlogHealthServer struct {
	pb.UnimplementedHealthServiceServer
	DB *gorm.DB
}

func NewBlogHealthServer(db *gorm.DB) *BlogHealthServer {
	return &BlogHealthServer{DB: db}
}

func (s *BlogHealthServer) Check(_ context.Context, _ *pb.CheckRequest) (*pb.CheckResponse, error) {
	return CheckService(s.DB)
}
