package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddArticlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/create-article.tmpl", gin.H{
		"title": "title",
	})
}

func AddArticle(c *gin.Context) {

}
