package router

import (
	"github.com/gin-gonic/gin"
	"gim/api"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	User := Router.Group("user")
	{
		User.POST("v1/login",api.Login)
		User.POST("v1/reg",api.Reg)
	}
}
