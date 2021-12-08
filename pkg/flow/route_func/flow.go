package route_func

import (
	"github.com/gin-gonic/gin"
	"yang-backend/pkg/command/command_func"
	command_models "yang-backend/pkg/command/models"
	"yang-backend/pkg/command/resp_code"
	"yang-backend/pkg/db"
	"yang-backend/pkg/flow/models"
)

type CreateFlowInput struct {
	Flow models.Flow `json:"flow"`
}

type CreateFlowOutput struct {
	command_models.ResponseBase
	Data models.Flow `json:"data"`
}

func CreateFlow(ctx *gin.Context) {
	var input CreateFlowInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&input.Flow).Error; err != nil {
		ctx.JSON(200, resp_code.CreateFlowFailed)
		return
	}
	output := CreateFlowOutput{
		ResponseBase: command_models.Success,
		Data:         input.Flow,
	}
	ctx.JSON(200, output)
}

type GetFlowsInput struct {
	command_models.GetListModel
}

type GetFlowsOutput struct {
	command_models.ResponseBase
	Data  []models.Flow `json:"data"`
	Total int64         `json:"total"`
}

func GetFlows(ctx *gin.Context) {
	var i GetFlowsInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(models.Flow{}, i.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.GetFlowsFailed)
		return
	}
	var flows []models.Flow
	if err := tx.Find(&flows).Error; err != nil {
		ctx.JSON(200, resp_code.GetFlowsFailed)
		return
	}
	o := GetFlowsOutput{
		ResponseBase: command_models.Success,
		Data: flows,
		Total: total,
	}
	ctx.JSON(200, o)
}
