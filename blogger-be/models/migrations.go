package models

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Blog{},
		&Share{},
		&Comment{},
		&Like{},
		&Tag{},
	)

}
