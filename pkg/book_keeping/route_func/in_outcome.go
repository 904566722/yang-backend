package route_func

import (
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/models"
    "yang-backend/pkg/command/command_func"
    command_models "yang-backend/pkg/command/models"
    "yang-backend/pkg/command/resp_code"
)

type GetCtgsInput struct {
    command_models.GetListModel
}
type GetCtgsOutput struct {
    command_models.ResponseBase
    Data models.CtgList `json:"data"`
    Total int64 `json:"total"`
}

// GetCtgs 获取分类列表
func GetCtgs(ctx *gin.Context) {
    var input GetCtgsInput
    if err := ctx.BindJSON(&input); err != nil {
        inputError := command_models.InputError(err)
        ctx.JSON(inputError.HttpCode(), inputError)
        return
    }
    var t1 int64
    tx, _, err := command_func.CommandGets(models.IncomeCategory{}, input.GetListModel)
    if err != nil {
        ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
        return
    }
    opUnit := ctx.Param("op_unit")
    var incomeCtgs []models.IncomeCategory
    if err := tx.Where("op_unit=?", opUnit).Find(&incomeCtgs).Error; err != nil {
        ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
        return
    }
    tx.Count(&t1)
    var t2 int64
    tx, _, err = command_func.CommandGets(models.OutcomeCategory{}, input.GetListModel)
    if err != nil {
        ctx.JSON(200, resp_code.OutcomeCategoryGetFailed)
        return
    }
    var outcomeCtgs []models.OutcomeCategory
    if err := tx.Where("op_unit=?", opUnit).Find(&outcomeCtgs).Error; err != nil {
        ctx.JSON(200, resp_code.OutcomeCategoryGetFailed)
        return
    }
    tx.Count(&t2)
    var ctgList models.CtgList
    ctgList.IncomeCtgs = incomeCtgs
    ctgList.OutcomeCtgs = outcomeCtgs
    output := GetCtgsOutput{
        ResponseBase: command_models.Success,
        Data: ctgList,
        Total: t1 + t2,
    }
    ctx.JSON(200, output)
}
