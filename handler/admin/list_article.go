package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/list-article.tmpl", gin.H{
		"title": "title",
	})
}
