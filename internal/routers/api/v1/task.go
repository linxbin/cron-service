package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/linxbin/corn-service/global"
	"github.com/linxbin/corn-service/internal/service"
	"github.com/linxbin/corn-service/pkg/app"
	"github.com/linxbin/corn-service/pkg/convert"
	"github.com/linxbin/corn-service/pkg/errcode"
)

type Task struct{}

func NewTask() Task {
	return Task{}
}

func (t Task) Create(c *gin.Context) {
	param := service.CreateTaskRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTask(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTaskFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t Task) Update(c *gin.Context) {
	params := service.UpDateTaskReuqest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&params)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTask err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTaskFail)
		return
	}

	response.ToResponse(gin.H{})

}
