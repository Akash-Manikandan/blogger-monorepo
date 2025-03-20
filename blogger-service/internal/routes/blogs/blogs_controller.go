package blogs

import (
	"context"

	"github.com/Akash-Manikandan/blogger-service/internal/utils"
	pb "github.com/Akash-Manikandan/blogger-service/proto/blog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type BlogServer struct {
	pb.UnimplementedBlogServiceServer
	DB *gorm.DB
}

func NewBlogServer(db *gorm.DB) *BlogServer {
	return &BlogServer{DB: db}
}

func (s *BlogServer) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	userID, ok := utils.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}
	return CreateBlogService(s.DB, req, userID)
}

func (s *BlogServer) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	userID, ok := utils.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}
	return GetBlogService(s.DB, req, userID)
}

func (s *BlogServer) ListBlogs(ctx context.Context, req *pb.ListBlogsRequest) (*pb.ListBlogsResponse, error) {
	userID, ok := utils.GetUserID(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}
	return ListBlogsService(s.DB, req, userID)
}
