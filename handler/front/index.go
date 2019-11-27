package front

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "front/index.tmpl", gin.H{
		"title": "Main website",
	})
}
