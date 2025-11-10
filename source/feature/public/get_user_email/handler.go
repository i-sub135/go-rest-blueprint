package get_user_email

import (
	customerrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/customer_repo"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Handler struct {
	repo Repositories
}

func NewHandler(userRepo *userrepo.UserRepo, customerRepo *customerrepo.CustomerRepo) *Handler {
	repo := injectRepository(userRepo, customerRepo)
	return &Handler{repo: repo}
}
