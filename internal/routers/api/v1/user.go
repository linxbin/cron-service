package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/corn-service/global"
	"github.com/linxbin/corn-service/internal/service"
	"github.com/linxbin/corn-service/pkg/app"
	"github.com/linxbin/corn-service/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u *User) Login(c *gin.Context) {
	param := service.UserLoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	userInfo, err := svc.Login(&param)
	if err != nil {
		global.Logger.Errorf("svc.user login err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserLoginFail)
		return
	}

	response.ToResponse(userInfo)
}
