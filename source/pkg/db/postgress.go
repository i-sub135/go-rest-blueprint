package db

import (
	"context"
	"time"

	"github.com/i-sub135/go-rest-blueprint/source/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	// Add connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := config.GetConfig().DB.DSN

	// Add connection timeout to DSN if not present
	if dsn[len(dsn)-1:] != "=" {
		dsn += " connect_timeout=10"
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Test connection with timeout
	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}

	// Ping with context timeout
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	return database, nil
}
