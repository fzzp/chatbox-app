package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域中间件
func EnableCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		// 预请求缓存时长
		MaxAge: time.Minute,
	})
}
