package service

import (
	getalluser "github.com/i-sub135/go-rest-blueprint/source/feature/public/get_all_user"

	"github.com/gin-gonic/gin"
	userrepo "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
	getuser "github.com/i-sub135/go-rest-blueprint/source/feature/public/get_user"
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
	userRoute := routeGroup.Group("/user")

	// userRoute.Use(middleware) uncommand for use middleware
	userRoute.GET("", getalluser.NewHandler(userRepo).Impl)
	userRoute.GET("/:id", getuser.NewHandler(userRepo).Impl)

}
