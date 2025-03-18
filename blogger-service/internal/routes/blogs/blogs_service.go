package blogs

import (
	"log"

	"github.com/Akash-Manikandan/blogger-service/models"
	pb "github.com/Akash-Manikandan/blogger-service/proto/blog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

func CreateBlogService(DB *gorm.DB, req *pb.CreateBlogRequest, userID string) (*pb.CreateBlogResponse, error) {
	blog := models.Blog{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
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
			Title:     blog.Title,
			Content:   blog.Content,
			Author:    blog.UserID,
			Id:        blog.ID,
			CreatedAt: blog.CreatedAt.String(),
			UpdatedAt: blog.UpdatedAt.String(),
			Views:     0,
			Likes:     0,
		},
	}, nil
}
