package router

import (
	"net/http"
	"zhangcs/blog/handler/front"
	"zhangcs/blog/handler/sd"
	"zhangcs/blog/handler/user"
	"zhangcs/blog/router/middleware"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/404")
	})

	g.GET("/", front.Index)
	g.GET("/detail/:id", front.Detail)
	g.GET("/404", front.NotFund)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}
	loadSd(g)
	g.POST("/login", user.Login)
	return g
}

func loadSd(g *gin.Engine) {
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
}
