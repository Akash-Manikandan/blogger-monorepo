package users

import (
	"context"
	"log"

	"github.com/Akash-Manikandan/blogger-be/internal/config"
	"github.com/Akash-Manikandan/blogger-be/internal/utils"
	"github.com/Akash-Manikandan/blogger-be/models"
	pb "github.com/Akash-Manikandan/blogger-be/proto/user/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer

	DB *gorm.DB
}

func UserServerRegister(db *gorm.DB) *UserServer {
	return &UserServer{DB: db}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.DB.Create(user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	token, err := config.GenerateToken(user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &pb.CreateUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Token:     token,
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	data, err := utils.VerifyPassword(user.Password, req.Password)

	if err != nil || !data {

		log.Println("Error verifying password:", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}
	token, err := config.GenerateToken(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &pb.LoginResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Token:     token,
	}, nil
}
