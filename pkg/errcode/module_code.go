package errcode

var (
	ErrorGetTaskListFail = NewError(20010001, "获取任务列表失败")
	ErrorCreateTaskFail  = NewError(20010002, "创建任务失败")
	ErrorUpdateTaskFail  = NewError(20010003, "更新任务失败")
	ErrorDeleteTaskFail  = NewError(20010004, "删除任务失败")
	ErrorCountTaskFail   = NewError(20010005, "统计任务失败")
)
