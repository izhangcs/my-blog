package router

import (
	"net/http"
	"zhangcs/blog/handler/admin"
	"zhangcs/blog/handler/front"
	"zhangcs/blog/handler/sd"
	"zhangcs/blog/router/middleware"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/404")
	})

	g.GET("/", front.Index)
	g.GET("/detail/:id", front.Detail)
	g.GET("/404", front.NotFund)

	g.GET("/e-login", admin.LoginPage)
	g.POST("/e-login", admin.Login)
	g.GET("/g", front.GeneratePassword)

	a := g.Group("/admin")
	//a.Use(middleware.AuthMiddleware())
	{
		a.GET("/articles/list", admin.ListArticle)
		// a.GET("/articles/add", admin.AddArticlePage)
		// a.POST("/articles/add", admin.AddArticle)
		// a.GET("/articles/edit", admin.EditArticlePage)
		// a.POST("/articles/edit", admin.EditArticle)
	}
	loadSd(g)
	return g
}

func loadAdmin(g *gin.Engine) *gin.Engine {
	a := g.Group("/admin")
	a.Use(middleware.AuthMiddleware())
	{
		a.GET("/articles/list", admin.ListArticle)
		a.GET("/articles/add", admin.AddArticlePage)
		a.POST("/articles/add", admin.AddArticle)
		a.GET("/articles/edit", admin.EditArticlePage)
		a.POST("/articles/edit", admin.EditArticle)
	}
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
