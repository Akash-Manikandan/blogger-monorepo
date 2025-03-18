package health

import (
	pb "github.com/Akash-Manikandan/blogger-be/proto/health/v1"
	"gorm.io/gorm"
)

func CheckService(DB *gorm.DB) (*pb.CheckResponse, error) {
	sqlDB, err := DB.DB()
	if err != nil {
		return &pb.CheckResponse{
			Status:   "OK",
			DbStatus: err.Error(),
		}, nil
	}

	err = sqlDB.Ping()
	if err != nil {
		return &pb.CheckResponse{
			Status:   "OK",
			DbStatus: err.Error(),
		}, nil
	}

	return &pb.CheckResponse{
		Status:   "OK",
		DbStatus: "OK",
	}, nil
}
