package front

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "front/detail.tmpl", gin.H{
		"title": "Main website",
	})
}
