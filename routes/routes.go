package routes

import (
	"acana/controller"
	"acana/logger"
	"acana/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	// rigister busniess router
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/comuunity/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)

		v1.GET("/posts/", controller.GetPostListHandler)
	}

	// r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "ok",
	// 	})
	// })

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
