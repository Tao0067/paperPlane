package http

import (
	"github.com/gin-gonic/gin"
	"go_xx/app/request"
	"go_xx/models"
	"go_xx/utils"
)

func CreateAircraft(ctx *gin.Context) {
	var params request.AircraftCreateRequest
	err := ctx.BindJSON(&params)
	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	aircraft := models.Aircraft{
		Name:         params.Name,
		Type:         params.Type,
		CoverPicture: params.CoverPicture,
		Content:      params.Content,
		Tag:          params.Tag,
	}
	err = aircraft.CreateAircraft()
	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	utils.SucceedMsg(ctx, "ok")
	return
}

func UpdateAircraft(ctx *gin.Context) {
	var params request.AircraftEditRequest
	err := ctx.BindJSON(&params)
	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}

	aircraft, err := models.FindAircraftById(params.Id)
	if err != nil {
		utils.Error(ctx, err.Error())
		return
	}
	aircraft.Name = params.Name

}
