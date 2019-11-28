package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.tmpl", gin.H{
		"title": "title",
	})
}
