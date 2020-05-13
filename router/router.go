package router

import (
	auth "cy/controller/auth"
	"cy/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()

	router.Use(middleware.RequestLog(), middleware.TranslationMiddleware(), middleware.RecoveryMiddleware())

	router.POST("register", auth.Register)
	router.POST("login", auth.Login)

	/**
	@desc 改分组下验证token
	*/
	authRization := router.Group("/", middleware.JWTAuth())
	{
		authRization.POST("/user-info", auth.UserInfo)
	}

	return router
}
