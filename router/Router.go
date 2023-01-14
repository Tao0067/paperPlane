package router

import (
	"github.com/gin-gonic/gin"
	"go_xx/common"
	"go_xx/handlers"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")

	v1.POST("/login", handlers.Login)

	user := v1.Group("adminUser")
	{
		user.POST("/add", handlers.AddAdminUser)
		user.GET("/detail", common.JWY(), handlers.DetailAdminUserById)
	}

	return r
}
