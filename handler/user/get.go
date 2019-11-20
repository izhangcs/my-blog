package user

import (
	. "zhangcs/blog/handler"
	"zhangcs/blog/model"
	"zhangcs/blog/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, map[string]interface{}{
		"username": user.Username,
	})
}
