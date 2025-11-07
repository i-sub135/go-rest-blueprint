package healtcheck

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i-sub135/go-rest-blueprint/source/config"
)

func (h *Handler) HealtCheck(c *gin.Context) {

	now := time.Now()
	if err := checkConnection(h.db, c.Request.Context()); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "FAIL",
			"error":   err.Error(),
			"version": config.GetConfig().App.Version,
			"time":    now,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"message":   "Database connection healthy",
		"version":   config.GetConfig().App.Version,
		"timestamp": now,
	})
}
