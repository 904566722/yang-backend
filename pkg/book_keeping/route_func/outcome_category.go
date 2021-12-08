package route_func

import (
	"github.com/gin-gonic/gin"
    "strings"
    "yang-backend/pkg/book_keeping/models"
    "yang-backend/pkg/command/command_func"
    command_models "yang-backend/pkg/command/models"
    "yang-backend/pkg/command/resp_code"
)

type GetOutcomeCtgsInput struct {
	command_models.GetListModel
	OpUnit string `json:"op_unit"`
}

type GetOutcomeCtgsOutput struct {
	command_models.ResponseBase
	Data  []models.OutcomeCategory `json:"data"`
	Total int64                    `json:"total"`
}

func GetOutcomeCategories(ctx *gin.Context) {
    var input GetOutcomeCtgsInput
    if err := ctx.BindJSON(&input); err != nil {
        inputError := command_models.InputError(err)
        ctx.JSON(inputError.HttpCode(), inputError)
        return
    }
    tx, total, err := command_func.CommandGets(models.OutcomeCategory{}, input.GetListModel)
    if err != nil {
        ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
        return
    }
    if input.OpUnit != "" && strings.Contains(input.OpUnit, "!") {
        opUnit := strings.Replace(input.OpUnit, "!", "", -1)
        tx.Where("op_unit != ?", opUnit)
    } else if input.OpUnit != "" {
        tx.Where("op_unit = ?", input.OpUnit)
    }
    var outcomeCtgs []models.OutcomeCategory
    if err := tx.Find(&outcomeCtgs).Error; err != nil {
        ctx.JSON(200, resp_code.OutcomeCategoryGetFailed)
        return
    }
    tx.Count(&total)
    output := GetOutcomeCtgsOutput{
        ResponseBase: command_models.Success,
        Data: outcomeCtgs,
        Total: total,
    }
    ctx.JSON(200, output)
}
