package router

import (
	"github.com/ardipermana59/mygram/internal/handler"
	"github.com/ardipermana59/mygram/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
		userRouter.PUT("/", middleware.AuthMiddleware(), handler.UpdateUser)
		userRouter.DELETE("/", middleware.AuthMiddleware(), handler.DeleteUser)
	}

	return r
}
