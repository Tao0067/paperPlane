package router

import (
	"github.com/gin-gonic/gin"
	"go_xx/app/http"
	"go_xx/utils"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")

	v1.POST("/admin/login", http.AdminLogin)

	user := v1.Group("adminUser")
	{
		user.POST("/add", http.AddAdminUser)
		user.GET("/detail", utils.JWY(), http.DetailAdminUserById)
	}

	return r
}
