package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	group := r.Group("/api/v1")
	{
		InitTaskRouter(group) // 任务管理
	}

	return r
}
