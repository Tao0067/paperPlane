package handlers

import (
	"github.com/gin-gonic/gin"
	"go_xx/common"
	"go_xx/models"
	"net/http"
)

type loginParams struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var params loginParams

	err := ctx.BindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	adminUser, err := models.FindAdminUserByNameAndPas(params.Name, params.Password)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	token, err := common.GenerateToken(adminUser.Id, "iss")

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": token,
	})
	return

}
