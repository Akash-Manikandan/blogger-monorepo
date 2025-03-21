package blogs

import (
	"errors"
	"log"

	"github.com/Akash-Manikandan/blogger-service/models"
	pb "github.com/Akash-Manikandan/blogger-service/proto/blog/v1"
	userPb "github.com/Akash-Manikandan/blogger-service/proto/user/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

func CreateBlogService(DB *gorm.DB, req *pb.CreateBlogRequest, userID string) (*pb.CreateBlogResponse, error) {
	blog := models.Blog{
		Title:    req.Title,
		Content:  req.Content,
		UserID:   userID,
		IsPublic: req.IsPublic,
	}
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(&blog).Error; err != nil {
			log.Printf("Error creating blog: %v", err)
			return err
		}
		if err := tx.Preload("User").Where("id = ?", blog.ID).First(&blog).Error; err != nil {
			log.Printf("Error fetching blog: %v", err)
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Errorf(codes.Canceled, "failed to create blog: %v", err)
	}

	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Title:   blog.Title,
			Content: blog.Content,
			Author: &userPb.User{
				Id:       blog.User.ID,
				Username: blog.User.Username,
				Email:    blog.User.Email,
			},
			Id:        blog.ID,
			CreatedAt: blog.CreatedAt.String(),
			UpdatedAt: blog.UpdatedAt.String(),
			IsPublic:  blog.IsPublic,
			Views:     int32(blog.ViewCount),
			Likes:     int32(blog.TrendingCount),
		},
	}, nil
}

func GetBlogService(DB *gorm.DB, req *pb.GetBlogRequest, userId string) (*pb.GetBlogResponse, error) {
	var blog models.Blog

	err := DB.
		Preload("User").
		Where("id = ? AND user_id = ?", req.Id, userId).
		Or("id = ? AND is_public = ?", req.Id, true).
		First(&blog).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "blog not found")
		}
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	return &pb.GetBlogResponse{
		Blog: &pb.Blog{
			Title:   blog.Title,
			Content: blog.Content,
			Author: &userPb.User{
				Id:       blog.User.ID,
				Username: blog.User.Username,
				Email:    blog.User.Email,
			},
			Id:        blog.ID,
			Views:     int32(blog.ViewCount),
			Likes:     int32(blog.TrendingCount),
			IsPublic:  blog.IsPublic,
			CreatedAt: blog.CreatedAt.String(),
			UpdatedAt: blog.UpdatedAt.String(),
		},
	}, nil
}

func ListBlogsService(DB *gorm.DB, req *pb.ListBlogsRequest, userId string) (*pb.ListBlogsResponse, error) {
	var blogs []models.Blog
	if err := DB.Preload("User").Where("user_id = ? or is_public = ?", userId, true).Find(&blogs).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "blogs not found: %v", err)
	}

	var blogResponses []*pb.BlogWithoutContent
	for _, blog := range blogs {
		blogResponses = append(blogResponses, &pb.BlogWithoutContent{
			Title: blog.Title,
			Author: &userPb.User{
				Id:       blog.User.ID,
				Username: blog.User.Username,
				Email:    blog.User.Email,
			},
			Id:        blog.ID,
			Views:     int32(blog.ViewCount),
			Likes:     int32(blog.TrendingCount),
			IsPublic:  blog.IsPublic,
			CreatedAt: blog.CreatedAt.String(),
			UpdatedAt: blog.UpdatedAt.String(),
		})
	}

	return &pb.ListBlogsResponse{
		Blogs: blogResponses,
	}, nil
}
