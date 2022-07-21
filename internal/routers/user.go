package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/linxbin/corn-service/internal/routers/api/v1"
)

// InitUserRouter 索引路由
func InitUserRouter(Router *gin.RouterGroup) {
	user := v1.NewUser()
	router := Router.Group("users")
	{
		router.POST("/login", user.Login) // 登录
	}
}
