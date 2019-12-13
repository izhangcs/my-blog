package middleware

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("username")
		if value, ok := isLogin.(string); ok && value != "" && value == viper.GetString("admin.username") {
			c.Next()
		}
		fmt.Println("test")
		c.Redirect(http.StatusFound, "/")
	}
}
