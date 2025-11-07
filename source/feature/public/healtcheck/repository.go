package healtcheck

import (
	"context"
	"time"

	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
	"gorm.io/gorm"
)

func checkConnection(db *gorm.DB, ctx context.Context) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var result int
	err := db.WithContext(ctx).Raw("SELECT 1").Scan(&result).Error
	if err != nil {
		logger.Error().Err(err).Caller().Msg("Database health check failed - query error")
		return err
	}

	if result != 1 {
		logger.Error().Int("result", result).Caller().Msg("Database health check failed - unexpected result")
		return gorm.ErrInvalidDB
	}

	return nil
}
