package front

import (
	. "zhangcs/blog/handler"
	"zhangcs/blog/pkg/auth"

	"github.com/gin-gonic/gin"
)

func GeneratePassword(c *gin.Context) {
	password := c.Query("password")

	pass, _ := auth.Encrypt(password)

	SendResponse(c, nil, gin.H{
		"password": password,
		"pass":     pass,
	})
}
