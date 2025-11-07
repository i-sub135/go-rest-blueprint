package healtcheck

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HealtCheck(c *gin.Context) {

	now := time.Now()
	if err := checkConnection(h.db, c.Request.Context()); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status": "FAIL",
			"error":  err.Error(),
			"time":   now,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"message":   "Database connection healthy",
		"timestamp": now,
	})
}
