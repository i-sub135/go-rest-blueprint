package httpresputils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i-sub135/go-rest-blueprint/source/config"
)

type response struct {
	Status     string    `json:"status"`
	Message    *string   `json:"message,omitempty"`
	Time       time.Time `json:"timestamp"`
	AppVersion string    `json:"app_version"`
	Data       *any      `json:"data,omitempty"`
}

var (
	cfg = config.GetConfig()
)

func HttpRespOK(c *gin.Context, data *any, msg *string) {
	c.JSON(http.StatusOK, response{
		Status:     http.StatusText(http.StatusOK),
		Time:       time.Now(),
		AppVersion: cfg.App.Version,
		Data:       data,
		Message:    msg,
	})
}
func HttpRespNotFound(c *gin.Context, msg *string) {
	c.JSON(http.StatusNotFound, response{
		Status:     http.StatusText(http.StatusNotFound),
		Message:    msg,
		AppVersion: cfg.App.Version,
		Time:       time.Now(),
	})
}

func HttpRespBadRequest(c *gin.Context, msg *string, err *error) {
	c.JSON(http.StatusBadRequest, response{
		Status:     http.StatusText(http.StatusBadRequest),
		Message:    msg,
		AppVersion: cfg.App.Version,
		Time:       time.Now(),
	})
}

func HttpRespBadGateway(c *gin.Context, msg *string, err *error) {
	c.JSON(http.StatusBadGateway, response{
		Status:     http.StatusText(http.StatusBadGateway),
		Message:    msg,
		AppVersion: cfg.App.Version,
		Time:       time.Now(),
	})
}
