package router

import (
	"cy/controller/authController"
	"cy/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.New()

	router.Use(middleware.RequestLog(), middleware.TranslationMiddleware(), middleware.RecoveryMiddleware())

	router.POST("register", authController.Register)
	router.POST("login", authController.Login)

	/**
	@desc 改分组下验证token
	*/
	authRization := router.Group("/", middleware.JWTAuth())
	{
		authRization.POST("/user-info", authController.UserInfo)
	}

	return router
}
