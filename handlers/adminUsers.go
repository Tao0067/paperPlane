package handlers

import (
	"github.com/gin-gonic/gin"
	"go_xx/models"
	"net/http"
	"strconv"
)

type createParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AddAdminUser(ctx *gin.Context) {
	var param createParams
	ctx.BindJSON(&param)

	adminUser := models.AdminUser{
		Name:     param.Name,
		Password: param.Password,
	}

	if err := adminUser.Create(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
	})
}

func DetailAdminUserById(ctx *gin.Context)  {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求参数错误",
		})
		return
	}

	adminUserId, _ := strconv.Atoi(id)
	adminUser, err := models.FindAdminUserById(adminUserId);

	if  err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data":  adminUser,
	})
}
