package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
			c.Request.Header.Set("X-Request-ID", requestID)
		}
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Set("request_id", requestID)
		c.Next()
	}
}
