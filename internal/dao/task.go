package dao

import (
	"strings"
	"time"

	"github.com/linxbin/corn-service/internal/model"
	"github.com/linxbin/corn-service/pkg/app"
)

type TaskForm struct {
	Id            uint32
	Name          string `binding:"Required;MaxSize(32)"`
	Spec          string
	Command       string `binding:"Required;MaxSize(256)"`
	Timeout       uint16 `binding:"Range(0,86400)"`
	RetryTimes    uint8
	RetryInterval uint16
	Remark        string
	Status        uint8 `binding:"oneof=0 1"`
}

func (d *Dao) CreateTask(form TaskForm) error {
	task := model.Task{
		Name:          form.Name,
		Spec:          form.Spec,
		Command:       strings.TrimSpace(form.Command),
		Timeout:       form.Timeout,
		RetryTimes:    form.RetryTimes,
		RetryInterval: form.RetryInterval,
		Status:        form.Status,
		Remark:        form.Remark,
		Model:         &model.Model{Created: time.Now(), Updated: time.Now()},
	}

	return task.Create(d.engine)
}

func (d *Dao) UpdateTask(id uint32, form TaskForm) error {
	task := model.Task{
		Model: &model.Model{ID: id},
	}
	values := model.CommonMap{
		"name":           form.Name,
		"spec":           form.Spec,
		"command":        form.Command,
		"timeout":        form.Timeout,
		"retry_times":    form.RetryTimes,
		"retry_interval": form.RetryInterval,
		"remark":         form.Remark,
		"status":         form.Status,
	}

	return task.Update(d.engine, values)
}

func (d *Dao) CountTask(name string, status uint8) (int, error) {
	task := model.Task{
		Name:   name,
		Status: status,
	}
	return task.Count(d.engine)
}

func (d *Dao) GetTaskList(name string, status uint8, page, pageSize int) ([]*model.Task, error) {
	task := model.Task{Name: name, Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return task.List(d.engine, pageOffset, pageSize)
}

// func (d *Dao) DeleteTag(id uint32) error {
// 	tag := model.Tag{Model: &model.Model{ID: id}}
// 	return tag.Delete(d.engine)
// }
