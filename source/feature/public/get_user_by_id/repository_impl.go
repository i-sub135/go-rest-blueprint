package get_user_by_id

import (
	"context"

	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
)

func (r *repositoryImpl) LogUserAccess(ctx context.Context, userID uint, requesterIP string) error {
	logger.Info().
		Uint("user_id", userID).
		Str("requester_ip", requesterIP).
		Msg("user profile accessed")

	return nil
}
