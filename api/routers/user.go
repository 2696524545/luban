package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/dnsjia/luban/pkg/core/user"
)

func UserRouter(r *gin.RouterGroup) {
	users := r.Group("user")
	{
		users.POST("/register", user.Register)
		users.POST("/login", user.Login)
	}
}
