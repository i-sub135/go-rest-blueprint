package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/i-sub135/go-rest-blueprint/source/service/constant"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(constant.RequestIDHeader)
		if strings.TrimSpace(reqID) == "" {
			reqID = uuid.New().String()[:8] // or uuid.New().String() for full UUID
		}
		c.Set(constant.RequestIDKey, reqID)
		c.Header(constant.RequestIDHeader, reqID)
		c.Next()
	}
}
