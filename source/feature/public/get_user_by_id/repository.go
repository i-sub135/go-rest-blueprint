package get_user_by_id

import (
	"context"

	usermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/user_model"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Repositories interface {
	// common repo implement
	GetByID(ctx context.Context, id uint) (*usermodel.User, error)

	// internal repo implement
	LogUserAccess(ctx context.Context, userID uint, requesterIP string) error
}

type repositoryImpl struct {
	*userrepo.UserRepo // Embedded shared repo
}

func injectRepository(userRepo *userrepo.UserRepo) Repositories {
	return &repositoryImpl{
		UserRepo: userRepo,
	}
}
