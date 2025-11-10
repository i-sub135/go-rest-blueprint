package get_all_user

import (
	"context"

	usermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/user_model"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Repositories interface {
	// common repo implement
	GetAll(ctx context.Context) (*[]usermodel.User, error)
	/**
	append methods name for implement internal methode
		example GetUserByEmail(ctx context.Context) (usermodel.User, error)
		and implement method/function in repository_impl.go

	*/
}

type repositoryImpl struct {
	*userrepo.UserRepo // Embedded shared repo
}

func injectRepository(userRepo *userrepo.UserRepo) Repositories {
	return &repositoryImpl{
		UserRepo: userRepo,
	}
}
