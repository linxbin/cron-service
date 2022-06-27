package service

import (
	"github.com/linxbin/corn-service/internal/dao"
	"github.com/linxbin/corn-service/internal/model"
	"github.com/linxbin/corn-service/pkg/app"
)

type CreateTaskRequest struct {
	TaskRequest
}

type UpDateTaskReuqest struct {
	ID uint32 `form:"id" binding:"required"`
	TaskRequest
}

type TaskRequest struct {
	Name          string `form:"name" binding:"required,min=0,max=32"`
	Spec          string `form:"spec" binding:"required,min=0,max=64"`
	Command       string `form:"command" binding:"required,min=0,max=255"`
	Timeout       uint16 `form:"timeout" binding:"required,gte=1,lte=86400"`
	RetryTimes    uint8  `form:"retryTimes" binding:"required,gte=0"`
	RetryInterval uint16 `form:"retryInterval" binding:"required,gte=1"`
	Remark        string `form:"remark" binding:"min=0,max=255"`
	Status        uint8  `form:"status" binding:"oneof=0 1"`
}

type TaskListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type CountTaskRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CreateTask(request *CreateTaskRequest) error {
	form := dao.TaskForm{
		Name:          request.Name,
		Spec:          request.Spec,
		Command:       request.Command,
		Timeout:       request.Timeout,
		RetryTimes:    request.RetryTimes,
		RetryInterval: request.RetryInterval,
		Remark:        request.Remark,
		Status:        request.Status,
	}

	return svc.dao.CreateTask(form)
}

func (svc *Service) UpdateTask(request *UpDateTaskReuqest) error {
	form := dao.TaskForm{
		Name:          request.Name,
		Spec:          request.Spec,
		Command:       request.Command,
		Timeout:       request.Timeout,
		RetryTimes:    request.RetryTimes,
		RetryInterval: request.RetryInterval,
		Remark:        request.Remark,
		Status:        request.Status,
	}

	return svc.dao.UpdateTask(request.ID, form)
}

func (svc *Service) CountTask(request *CountTaskRequest) (int, error) {
	return svc.dao.CountTask(request.Name, request.Status)
}

func (svc *Service) GetTaskList(request *TaskListRequest, pager *app.Pager) ([]*model.Task, error) {
	return svc.dao.GetTaskList(request.Name, request.Status, pager.Page, pager.PageSize)
}
