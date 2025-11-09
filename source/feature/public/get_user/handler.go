package getuser

import (
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Handler struct {
	repo UserRepositories
}

func NewHandler(userRepo *userrepo.UserRepo) *Handler {
	repo := injectRepository(userRepo)
	return &Handler{repo: repo}
}
