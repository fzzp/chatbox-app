package routes

import (
	v1 "chatbox-app/api/v1"
	"chatbox-app/middleware"
	"chatbox-app/socket"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	authApi v1.AuthApi
	ss      socket.SocketServer
)

func NewApiRoutes() http.Handler {
	mux := gin.New()
	mux.Use(gin.Logger())
	mux.Use(middleware.Recovery())
	mux.Use(middleware.NewRequestID())
	mux.Use(middleware.EnableCORS())

	r1 := mux.Group("/api/v1")
	{
		r1.POST("/signup", authApi.SignUp)
		r1.POST("/login", authApi.Login)
		r1.POST("/logout", authApi.Logout)
	}

	mux.GET("/ws", ss.WsEndpoint)

	// swagger docs
	mux.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return mux
}
