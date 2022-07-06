package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/linxbin/corn-service/internal/routers/api/v1"
)

// InitTaskRouter 索引路由
func InitTaskRouter(Router *gin.RouterGroup) {

	task := v1.NewTask()
	router := Router.Group("tasks")
	{
		router.POST("", task.Create)       // 创建任务
		router.PUT("/:id", task.Update)    // 更新任务
		router.DELETE("/:id", task.Delete) // 删除任务
		router.GET("", task.List)          // 任务列表
		router.GET("/:id", task.Detail)    // 任务详情
	}
}
