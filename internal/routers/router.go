package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	group := r.Group("/api/v1")
	{
		InitTaskRouter(group)    // 任务管理
		InitTaskLogRouter(group) // 任务日志管理
		InitUserRouter(group)    // 用户管理
	}

	return r
}
