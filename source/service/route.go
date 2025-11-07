package service

import (
	"github.com/gin-gonic/gin"
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

}
