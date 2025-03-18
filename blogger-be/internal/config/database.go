package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global DB connection
func InitDB(cfg *Config) (*gorm.DB, error) {
	if cfg.DATABASE_URL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// Open PostgreSQL connection with SSL enabled
	db, err := gorm.Open(postgres.Open(cfg.DATABASE_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Reduce log verbosity for performance
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get database instance:", err)
	}

	// Set connection pooling and performance optimizations
	sqlDB.SetMaxOpenConns(25) // Maximum open connections
	sqlDB.SetMaxIdleConns(10) // Maximum idle connections
	// sqlDB.SetConnMaxLifetime(5 * 60 * 1000000000) // 5 minutes

	log.Println("✅ Database connected successfully!")

	return db, nil
}
