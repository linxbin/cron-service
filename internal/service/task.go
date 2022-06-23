package service

import (
	"github.com/linxbin/corn-service/internal/dao"
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

func (svc *Service) UpdateTag(request *UpDateTaskReuqest) error {
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

	return svc.dao.UpdateTag(request.ID, form)
}
