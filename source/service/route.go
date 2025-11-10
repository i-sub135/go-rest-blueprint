package service

import (
	"github.com/i-sub135/go-rest-blueprint/source/feature/public/get_all_user"
	"github.com/i-sub135/go-rest-blueprint/source/feature/public/get_user_by_id"
	"github.com/i-sub135/go-rest-blueprint/source/feature/public/get_user_email"

	"github.com/gin-gonic/gin"
	customerrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/customer_repo"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"

	"gorm.io/gorm"
)

type Routers struct {
	db *gorm.DB
}

func NewRouters(db *gorm.DB) *Routers {
	return &Routers{
		db: db,
	}
}

func (r *Routers) MountRouters(routeGroup *gin.RouterGroup) {

	// endpoint group user
	userRepo := userrepo.NewUserRepo(r.db)
	custRepo := customerrepo.NewRepo(r.db)
	userRoute := routeGroup.Group("/users")

	// userRoute.Use(middleware) uncommand for use middleware
	userRoute.GET("", get_all_user.NewHandler(userRepo))
	userRoute.GET("/:id", get_user_by_id.NewHandler(userRepo))
	userRoute.GET("/email", get_user_email.NewHandler(userRepo, custRepo))

}
