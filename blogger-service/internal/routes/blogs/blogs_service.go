package blogs

import (
	"errors"
	"log"
	"math"

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

	baseQuery := DB.Preload("User").Where("user_id = ? OR is_public = ?", userId, true)

	if req.Page != nil {
		var limit int32 = 100
		if req.Limit != nil {
			limit = req.GetLimit()
		}

		currentPage := req.GetPage()
		offset := (currentPage - 1) * limit

		var totalCount int64
		if err := baseQuery.Model(&models.Blog{}).Count(&totalCount).Error; err != nil {
			return nil, status.Errorf(codes.Internal, "failed to count blogs: %v", err)
		}

		if err := baseQuery.Order("updated_at DESC").Offset(int(offset)).Limit(int(limit)).Find(&blogs).Error; err != nil {
			return nil, status.Errorf(codes.NotFound, "blogs not found: %v", err)
		}

		totalPages := int32(math.Ceil(float64(totalCount) / float64(limit)))

		var nextPage, prevPage int32
		if currentPage < totalPages {
			nextPage = currentPage + 1
		}
		if currentPage > 1 {
			prevPage = currentPage - 1
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
			Blogs:       blogResponses,
			TotalCount:  int32(totalCount),
			CurrentPage: currentPage,
			TotalPages:  totalPages,
			NextPage:    nextPage,
			PrevPage:    prevPage,
		}, nil
	}

	if err := baseQuery.Order("updated_at DESC").Find(&blogs).Error; err != nil {
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

func UpdateBlogService(DB *gorm.DB, req *pb.UpdateBlogRequest, userId string) (*pb.UpdateBlogResponse, error) {
	var blog models.Blog
	err := DB.Preload("User").Where("id = ? AND user_id = ?", req.Id, userId).First(&blog).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "blog not found")
		}
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	blog.Title = req.GetTitle()
	blog.Content = req.GetContent()
	blog.IsPublic = req.GetIsPublic()

	err = DB.Save(&blog).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	return &pb.UpdateBlogResponse{
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

func DeleteBlogService(DB *gorm.DB, req *pb.DeleteBlogRequest, userId string) (*pb.DeleteBlogResponse, error) {
	var blog models.Blog
	err := DB.Preload("User").Where("id = ? AND user_id = ?", req.Id, userId).First(&blog).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "blog not found")
		}
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	err = DB.Delete(&blog).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	return &pb.DeleteBlogResponse{
		Message: "Blog deleted successfully",
	}, nil
}
