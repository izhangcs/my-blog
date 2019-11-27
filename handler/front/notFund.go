package front

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFund(c *gin.Context) {
	c.HTML(http.StatusOK, "front/404.tmpl", gin.H{
		"title": "404 Not Found",
	})
}
