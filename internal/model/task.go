package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	*Model
	Name          string `json:"name"`
	Spec          string `json:"spec"`
	Command       string `json:"command"`
	Timeout       uint16 `json:"timeout"`
	RetryTimes    uint8  `json:"retry_times"`
	RetryInterval uint16 `json:"retry_interval"`
	Remark        string `json:"remark"`
	Status        uint8  `json:"status"`
}

const (
	TaskStatusActive = 1
)

func (t Task) TableName() string {
	return "task"
}

func (t Task) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Task) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, NoDelete).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Task) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, NoDelete).Delete(&t).Error
}

func (t Task) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("status = ?", t.Status)
	if err := db.Model(&t).Where("is_del = ?", NoDelete).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Task) List(db *gorm.DB, pageOffset, pageSize int) ([]*Task, error) {
	var tasks []*Task
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("status = ?", t.Status)

	if err = db.Where("is_del = ?", NoDelete).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t Task) Detail(db *gorm.DB, ID uint32) (Task, error) {
	task := Task{}
	var err error

	if err = db.First(&task, "id = ? and is_del = ?", ID, NoDelete).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (t Task) ActiveList(db *gorm.DB, pageOffset, pageSize int) ([]*Task, error) {
	var tasks []*Task
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("status = ?", TaskStatusActive)

	if err = db.Where("is_del = ?", NoDelete).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
