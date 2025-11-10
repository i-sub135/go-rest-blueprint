package get_user_by_id

import (
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Handler struct {
	repo Repositories
}

func NewHandler(userRepo *userrepo.UserRepo) *Handler {
	repo := injectRepository(userRepo)
	return &Handler{repo: repo}
}
