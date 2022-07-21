package errcode

var (
	ErrorGetTaskListFail   = NewError(20010001, "获取任务列表失败")
	ErrorCreateTaskFail    = NewError(20010002, "创建任务失败")
	ErrorUpdateTaskFail    = NewError(20010003, "更新任务失败")
	ErrorDeleteTaskFail    = NewError(20010004, "删除任务失败")
	ErrorCountTaskFail     = NewError(20010005, "统计任务失败")
	ErrorTaskNotFound      = NewError(20010006, "任务不存在")
	ErrorTaskLogListFail   = NewError(20010011, "获取任务日志列表失败")
	ErrorTaskLogDetailFail = NewError(20010012, "获取任务日志详情失败")
	ErrorUserLoginFail     = NewError(20000000, "账号或者密码错误")
)
