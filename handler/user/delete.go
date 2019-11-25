package user

import (
	"strconv"
	. "zhangcs/blog/handler"
	"zhangcs/blog/model"
	"zhangcs/blog/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary 删除用户
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user/:id [DELETE]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
