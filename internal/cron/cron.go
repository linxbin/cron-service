package cron

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jakecoffman/cron"
	"github.com/linxbin/corn-service/global"
	"github.com/linxbin/corn-service/internal/dao"
	"github.com/linxbin/corn-service/internal/model"
)

type Cron struct {
	dao *dao.Dao
}

// TaskCount 任务计数
type TaskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func (tc *TaskCount) Add() {
	tc.wg.Add(1)
}

func (tc *TaskCount) Done() {
	tc.wg.Done()
}

func (tc *TaskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *TaskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

func NewCron() Cron {
	return Cron{
		dao: dao.New(global.DBEngine),
	}
}

var (
	serviceCron *cron.Cron
	taskCount   TaskCount // 任务计数-正在运行的任务
)

func (c Cron) Initialize() error {
	serviceCron = cron.New()
	serviceCron.Start()
	page := 1
	pageSize := 10
	taskNum := 0
	for {
		taskList, err := c.dao.TaskActiveList(page, pageSize)
		if err != nil {
			return err
		}
		if len(taskList) == 0 {
			break
		}
		for _, item := range taskList {
			if err = addTask(*item); err != nil {
				return err
			}
			taskNum++
		}
		page++
	}
	return nil
}

func addTask(task model.Task) error {
	taskFunc := createJob(task)
	if taskFunc == nil {
		return errors.New("创建任务处理Job失败")
	}

	cronName := strconv.Itoa(int(task.ID))
	err := panicToError(func() {
		serviceCron.AddFunc(task.Spec, taskFunc, cronName)
	})
	if err != nil {
		global.Logger.Errorf("添加任务到调度器失败: %s", err)
		return err
	}

	fmt.Printf(task.Name)
	return nil
}

func createJob(task model.Task) cron.FuncJob {
	taskFunc := func() {
		taskCount.Add()
		defer taskCount.Done()

		taskLogId, err := beforeExecJob(task)
		if err != nil || taskLogId <= 0 {
			return
		}
		global.Logger.Infof("开始执行任务#%s#命令-%s", task.Name, task.Command)
		result := execJob(task)
		global.Logger.Infof("任务完成#%s#命令-%s", task.Name, task.Command)
		if err = afterExecJob(taskLogId, result); err != nil {
			return
		}
	}

	return taskFunc
}

func beforeExecJob(task model.Task) (taskLogId uint32, err error) {
	taskLogId, err = createTaskLog(task)
	if err != nil {
		return 0, err
	}

	return taskLogId, nil
}

func createTaskLog(task model.Task) (id uint32, err error) {
	taskLogForm := dao.TaskLogForm{
		TaskId:     task.ID,
		Name:       task.Name,
		Spec:       task.Spec,
		Command:    task.Command,
		Timeout:    task.Timeout,
		RetryTimes: task.RetryTimes,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Status:     model.TaskLogStatusPending,
	}
	d := dao.New(global.DBEngine)
	return d.CreateTaskLog(taskLogForm)
}

func updateTaskLog(taskLogId uint32, result TaskResult) error {
	d := dao.New(global.DBEngine)
	var status int
	if result.Err != nil {
		status = model.TaskLogStatusFailure
	} else {
		status = model.TaskLogStatusComplete
	}
	values := model.CommonMap{
		"status":      status,
		"end_time":    time.Now(),
		"retry_times": result.RetryTimes,
		"result":      result.Result,
	}
	return d.UpdateTaskLog(taskLogId, values)
}

// 任务执行后置操作
func afterExecJob(taskLogId uint32, result TaskResult) error {
	return updateTaskLog(taskLogId, result)
}

type TaskResult struct {
	Result     string
	Err        error
	RetryTimes uint8
}

// 执行具体任务
func execJob(task model.Task) TaskResult {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Errorf("panic#service/task.go:execJob#%s", err)
		}
	}()
	// 默认只运行任务一次
	var execTimes uint8 = 1
	if task.RetryTimes > 0 {
		execTimes += task.RetryTimes
	}
	var i uint8 = 0
	var output string
	var err error
	var cmd *exec.Cmd
	var cmdResult []byte
	for i < execTimes {
		cmd = exec.Command(
			"cmd",
			"/c",
			task.Command)
		if cmdResult, err = cmd.Output(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(cmdResult))
		// 指定参数后过滤换行符
		fmt.Println(strings.Trim(string(cmdResult), "\n"))
		i++
		if i < execTimes {
			global.Logger.Infof("任务执行失败#任务id-%d#重试第%d次#输出-%s#错误-%s", task.ID, i, output, err.Error())
			if task.RetryInterval > 0 {
				time.Sleep(time.Duration(task.RetryInterval) * time.Second)
			} else {
				// 默认重试间隔时间，每次递增1分钟
				time.Sleep(time.Duration(i) * time.Minute)
			}
		}
	}

	return TaskResult{Result: output, Err: err, RetryTimes: task.RetryTimes}
}

// PanicToError Panic转换为error
func panicToError(f func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(panicTrace(e))
		}
	}()
	f()
	return
}

// PanicTrace panic调用链跟踪
func panicTrace(err interface{}) string {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)

	return fmt.Sprintf("panic: %v %s", err, stackBuf[:n])
}
