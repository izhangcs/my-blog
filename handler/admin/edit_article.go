package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditArticlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/edit-article.tmpl", gin.H{
		"title": "ahgaha",
	})
}

func EditArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/edit-article.tmpl", gin.H{
		"title": "ahgaha",
	})
}
