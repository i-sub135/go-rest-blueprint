package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/i-sub135/go-rest-blueprint/source/service/constant"
)

func generateReqID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(constant.RequestIDHeader)
		if strings.TrimSpace(reqID) == "" {
			reqID = generateReqID()
		}
		c.Set(constant.RequestIDKey, reqID)
		c.Header(constant.RequestIDHeader, reqID)
		c.Next()
	}
}
