package model

import "github.com/jinzhu/gorm"

type Task struct {
	*Model
	Name          string `json:"name"`
	Spec          string `json:"spec"`
	Command       string `json:"command"`
	Timeout       uint16 `json:"timeout"`
	RetryTimes    uint8  `json:"retryTimes"`
	RetryInterval uint16 `json:"retryInterval"`
	Remark        string `json:"remark"`
	Status        uint8  `json:"status"`
}

func (t Task) TableName() string {
	return "task"
}

func (t Task) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Task) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ?", t.ID).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Task) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
