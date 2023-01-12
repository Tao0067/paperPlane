package router

import (
	"github.com/gin-gonic/gin"
	"go_xx/handlers"
)

func Router() *gin.Engine  {
	r := gin.Default()
	v1 := r.Group("v1")

	user := v1.Group("adminUser")
	{
		user.POST("/add", handlers.AddAdminUser)
		user.GET("/detail", handlers.DetailAdminUserById)
	}

	return r
}
