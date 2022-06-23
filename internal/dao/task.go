package dao

import (
	"strings"
	"time"

	"github.com/linxbin/corn-service/internal/model"
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

// func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
// 	tag := model.Task{
// 		Model: &model.Model{ID: id},
// 	}
// 	values := map[string]interface{}{
// 		"state":       state,
// 		"modified_by": modifiedBy,
// 	}
// 	if name != "" {
// 		values["name"] = name
// 	}

// 	return tag.Update(d.engine, values)
// }

// func (d *Dao) DeleteTag(id uint32) error {
// 	tag := model.Tag{Model: &model.Model{ID: id}}
// 	return tag.Delete(d.engine)
// }
