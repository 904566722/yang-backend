package route_func

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yang-backend/pkg/book_keeping/models"
	"yang-backend/pkg/command/command_func"
	command_models "yang-backend/pkg/command/models"
	"yang-backend/pkg/command/resp_code"
	"yang-backend/pkg/db"
)

type CreateIncomeInput struct {
	Income models.Income `json:"income"`
}

type CreateIncomeResponse struct {
	command_models.ResponseBase
	Data models.Income `json:"data"`
}

// IncomeCreate 新增收入条目
func IncomeCreate(ctx *gin.Context) {
	var input CreateIncomeInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&input.Income).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeCreateFailed)
		return
	}
	resp := CreateIncomeResponse{
		ResponseBase: command_models.Success,
		Data: input.Income,
	}
	ctx.JSON(200, resp)
}

type GetIncomesInput struct {
	command_models.GetListModel
}

type GetIncomesOutput struct {
	command_models.ResponseBase
	Data  []models.Income `json:"data"`
	Total int64           `json:"total"`
}

type GetIncomeOutput struct {
	command_models.ResponseBase
	Data  models.Income `json:"data"`
}

// IncomeGet 获取单个收入
func IncomeGet(ctx *gin.Context)  {
	incomeId := ctx.Param("income_id")
	var income models.Income
	if err := db.DB.Preload("IncomeCategory").Where("id = ?", incomeId).First(&income).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeGetFailed)
		return
	}
	output := GetIncomeOutput{
		ResponseBase: command_models.Success,
		Data: income,
	}
	ctx.JSON(200, output)
}

func getIncomeById(id string) ([]models.Income, error) {
	var incomes []models.Income
	err := db.DB.Where("id = ?", id).Find(&incomes).Error
	return incomes, err
}

type IncomeUpdateInput struct {
	Income models.Income `json:"income"`
}

type IncomeUpdateOutput struct {
	command_models.ResponseBase
	Income models.Income `json:"income"`
}

// IncomeUpdate 修改收入条目信息
func IncomeUpdate(ctx *gin.Context)  {
	incomes, err := getIncomeById(ctx.Param("income_id"))
	if err == nil && len(incomes) == 0 {
		resp_code.NotFoundResource.Message = fmt.Sprintf("未找到收入条目, income_id: %s", ctx.Param("income_id"))
		ctx.JSON(200, resp_code.NotFoundResource)
		return
	} else if err != nil {
		ctx.JSON(200, resp_code.IncomeGetFailed)
		return
	}
	oldIncome := incomes[0]
	var input IncomeUpdateInput
	input.Income = oldIncome
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Save(&input.Income).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeUpdateFailed)
		return
	}
	output := IncomeUpdateOutput{
		ResponseBase: command_models.Success,
		Income: input.Income,
	}
	ctx.JSON(200, output)
}

func IncomeDelete(ctx *gin.Context)  {
	incomeId := ctx.Param("income_id")
	if err := db.DB.Where("id = ?", incomeId).Delete(&models.Income{}).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeDeleteFailed)
		return
	}
	ctx.JSON(200, command_models.Success)
}

// IncomeGets 获取收入列表
func IncomeGets(ctx *gin.Context) {
	var input GetIncomesInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(&models.Income{}, input.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.IncomeGetsFailed)
		return
	}

	var incomes []models.Income
	// 查询
	if err := tx.Find(&incomes).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeGetsFailed)
		return
	}
	// 返回结果
	resp := GetIncomesOutput{
		ResponseBase: command_models.Success,
		Data:         incomes,
		Total:        total,
	}
	ctx.JSON(200, resp)
}
