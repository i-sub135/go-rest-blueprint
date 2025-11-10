package get_user_email

import (
	"github.com/gin-gonic/gin"
	customerrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/customer_repo"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
)

type Handler struct {
	repo Repositories
}

func NewHandler(userRepo *userrepo.UserRepo, customerRepo *customerrepo.CustomerRepo) gin.HandlerFunc {
	repo := injectRepository(userRepo, customerRepo)
	handler := Handler{repo: repo}
	return handler.Impl
}
