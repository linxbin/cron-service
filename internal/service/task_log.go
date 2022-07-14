package service

import (
	"github.com/linxbin/corn-service/internal/model"
	"github.com/linxbin/corn-service/pkg/app"
)

type TaskLogListRequest struct {
	TaskId uint32 `form:"task_id" binding:"required,gte=1"`
}

type CountTaskLogRequest struct {
	TaskId uint32 `form:"task_id" binding:"required,gte=1"`
}

func (svc *Service) CountTaskLog(request *CountTaskLogRequest) (int, error) {
	return svc.dao.CountTaskLog(request.TaskId)
}

func (svc *Service) TaskLogList(request *TaskLogListRequest, pager *app.Pager) ([]*model.TaskLog, error) {
	return svc.dao.TaskLogList(request.TaskId, pager.Page, pager.PageSize)
}
