package admin

import (
	"fmt"
	"net/http"
	. "zhangcs/blog/handler"
	"zhangcs/blog/pkg/auth"
	"zhangcs/blog/pkg/errno"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.tmpl", gin.H{
		"title": "title",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	isChecked := c.PostForm("is_checked")
	fmt.Println(isChecked)
	admin := viper.GetString("admin.username")
	passHash := viper.GetString("admin.password")
	if admin != username {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	if err := auth.Compare(passHash, password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	// 存储session
	session := sessions.Default(c)
	session.Set("username", username)
	session.Save()
	c.Redirect(http.StatusFound, "/admin/articles/list")
}
