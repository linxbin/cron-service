package dao

import (
	"github.com/linxbin/corn-service/internal/model"
	"strings"
	"time"
)

type TaskLogForm struct {
	Id         uint32
	TaskId     uint32
	Name       string `binding:"Required;MaxSize(32)"`
	Spec       string
	Command    string `binding:"Required;MaxSize(256)"`
	Timeout    uint16 `binding:"Range(0,86400)"`
	RetryTimes uint8
	Status     uint8 `binding:"oneof=0 1 2"`
	StartTime  time.Time
	EndTime    time.Time
	Result     string
}

func (d *Dao) CreateTaskLog(form TaskLogForm) (uint32, error) {
	taskLog := model.TaskLog{
		Name:       form.Name,
		Spec:       form.Spec,
		Command:    strings.TrimSpace(form.Command),
		Timeout:    form.Timeout,
		RetryTimes: form.RetryTimes,
		Status:     form.Status,
		StartTime:  form.StartTime,
		EndTime:    form.EndTime,
		Result:     form.Result,
		Model:      &model.Model{Created: time.Now(), Updated: time.Now()},
	}

	err := taskLog.Create(d.engine)
	if err != nil {
		return 0, err
	}

	return taskLog.ID, nil
}

func (d *Dao) UpdateTaskLog(id uint32, commonMap model.CommonMap) error {
	taskLog := model.TaskLog{
		Model: &model.Model{ID: id},
	}

	return taskLog.Update(d.engine, commonMap)
}
