package http

import (
	"github.com/gin-gonic/gin"
	"go_xx/app/request"
	"go_xx/models"
	"go_xx/utils"
	"net/http"
)

type loginParams struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AdminLogin(ctx *gin.Context) {
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

	token, err := utils.GenerateToken(adminUser.Id, "iss")

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

//Register 注册
func Register(ctx *gin.Context) {
	var registerRequest request.RegisterRequest
	err := ctx.BindJSON(registerRequest)

	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	user, err := models.FindUserByNameOrEmail(registerRequest.Name, registerRequest.Email)
	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	if user.Id > 0 {
		utils.Error(ctx, "用户名或者邮箱已经使用了！")
		return
	}

	userModel := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	if err := userModel.CreateUser(); err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	utils.SucceedMsg(ctx, "创建成功")
}

// Login 登陆
func Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	err := ctx.BindJSON(loginRequest)

	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	userModel, err := models.FindUserByEmailAndPassword(loginRequest.Email, loginRequest.Password)

	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	token, err := utils.GenerateToken(userModel.Id, "iss")
	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	utils.SucceedData(ctx, token)
}
