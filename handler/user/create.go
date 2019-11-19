package user

import (
	"fmt"
	"net/http"
	"zhangcs/blog/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(ctx *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := ctx.Bind(&r); err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
	}
	log.Debugf("username is: [%s], password is: [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	code, message := errno.DecodeErr(err)
	ctx.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
