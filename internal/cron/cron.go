package cron

import (
	"fmt"
	"github.com/linxbin/corn-service/global"
	"github.com/linxbin/corn-service/internal/dao"
	"github.com/linxbin/corn-service/internal/model"
	"github.com/robfig/cron/v3"
	"sync"
)

type Cron struct {
	dao *dao.Dao
}

// TaskCount 任务计数
type TaskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func NewCron() Cron {
	return Cron{
		dao: dao.New(global.DBEngine),
	}
}

var (
	serviceCron *cron.Cron
)

func (c Cron) Initialize() error {
	serviceCron := cron.New()
	serviceCron.Start()
	global.Logger.Infof("开始初始化定时任务")
	page := 1
	pageSize := 10
	for {
		taskList, err := c.dao.GetTaskActiveList(page, pageSize)
		if err != nil {
			return err
		}
		if len(taskList) == 0 {
			break
		}
		for _, item := range taskList {
			if err = c.AddTask(item); err != nil {
				return err
			}
		}
		page++
	}
	global.Logger.Infof("定时任务初始化完成, 共%d个定时任务添加到调度器")
	return nil
}

func (c Cron) AddTask(task *model.Task) error {
	fmt.Printf(task.Name)
	return nil
}
