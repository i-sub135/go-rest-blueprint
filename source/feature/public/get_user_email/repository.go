package get_user_email

import (
	"context"

	customermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/customer_model"
	usermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/user_model"
	customerrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/customer_repo"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Repositories interface {
	// common repo implement
	GetByEmail(ctx context.Context, email string) (*usermodel.User, error)
	GetCustomerFirstName(ctx context.Context, name string) (customer *[]customermodel.Customer, err error)

	// internal repo implement
}

type repositoryImpl struct {
	*userrepo.UserRepo         // Embedded user repo
	*customerrepo.CustomerRepo // Embedded customer repo
}

func injectRepository(userRepo *userrepo.UserRepo, customerRepo *customerrepo.CustomerRepo) Repositories {
	return &repositoryImpl{
		UserRepo:     userRepo,
		CustomerRepo: customerRepo,
	}
}
