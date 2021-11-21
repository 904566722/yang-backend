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

type CreateIncomeCategoryInput struct {
	IncomeCategory models.IncomeCategory `json:"income_category"`
}

type CreateIncomeCategoryResponse struct {
	command_models.ResponseBase
	Data models.IncomeCategory `json:"data"`
}

func IncomeCategoryCreate(ctx *gin.Context) {
	var input CreateIncomeCategoryInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&input.IncomeCategory).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeCategoryCreateFailed)
		return
	}
	resp := CreateIncomeCategoryResponse{ResponseBase: command_models.Success}
	resp.Data = input.IncomeCategory
	ctx.JSON(200, resp)
}

type IncomeCategoryUpdateInput struct {
	IncomeCategory models.IncomeCategory `json:"income_category"`
}

type IncomeCategoryUpdateOutput struct {
	command_models.ResponseBase
	IncomeCategory models.IncomeCategory `json:"income_category"`
}

func IncomeCategoryUpdate(ctx *gin.Context) {
	id := ctx.Param("income_category_id")
	incomeCategories, err := getIncomeCategoryById(id)
	if err == nil && len(incomeCategories) == 0 {
		resp_code.NotFoundResource.Message = fmt.Sprintf("未找到收入类别, income_category_id: %s", id)
		ctx.JSON(200, resp_code.NotFoundResource)
		return
	} else if err != nil {
		resp_code.IncomeCategoryGetFailed.Err = err
		ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
		return
	}

	var input IncomeCategoryUpdateInput
	input.IncomeCategory = incomeCategories[0]
	if err = ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}

	if err := db.DB.Save(&input.IncomeCategory).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeCategoryUpdateFailed)
		return
	}
	output := IncomeCategoryUpdateOutput{
		ResponseBase:   command_models.Success,
		IncomeCategory: input.IncomeCategory,
	}
	ctx.JSON(200, output)
}

func getIncomeCategoryById(id string) ([]models.IncomeCategory, error) {
	var incomeCategories []models.IncomeCategory
	err := db.DB.Where("id = ?", id).Find(&incomeCategories).Error
	return incomeCategories, err
}

type IncomeCategoryGetOutput struct {
	command_models.ResponseBase
	IncomeCategory models.IncomeCategory `json:"income_category"`
	IncomesNum     int64                 `json:"incomes_num"`
}

func IncomeCategoryGet(ctx *gin.Context) {
	id := ctx.Param("income_category_id")
	showIncomesNum := ctx.Query("show_incomes_num")
	incomeCategories, err := getIncomeCategoryById(id)
	if err == nil && len(incomeCategories) == 0 {
		resp_code.NotFoundResource.Message = fmt.Sprintf("未找到收入类别, income_category_id: %s", id)
		ctx.JSON(200, resp_code.NotFoundResource)
		return
	} else if err != nil {
		resp_code.IncomeCategoryGetFailed.Err = err
		ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
		return
	}
	output := IncomeCategoryGetOutput{
		ResponseBase:   command_models.Success,
		IncomeCategory: incomeCategories[0],
	}
	if showIncomesNum == "true" {
		var incomesNum int64
		if err := db.DB.Model(&models.Income{IncomeCategoryID: id}).Count(&incomesNum).Error; err != nil {
			resp_code.IncomeCountFailed.Message = "统计收入类别的收入条目出错"
			ctx.JSON(200, resp_code.IncomeCountFailed)
			return
		}
		output.IncomesNum = incomesNum
	}
	ctx.JSON(200, output)
}

type IncomeCategoryGetsInput struct {
	command_models.GetListModel
}

type IncomeCategoryGetsOutput struct {
	command_models.ResponseBase
	IncomeCategories []models.IncomeCategory `json:"income_categories"`
	Total            int64                   `json:"total"`
}

func IncomeCategoryGets(ctx *gin.Context) {
	var input IncomeCategoryGetsInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	// 通用方法
	tx, total, err := command_func.CommandGets(models.IncomeCategory{}, input.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
		return
	}

	var incomeCategories []models.IncomeCategory
	if err := tx.Find(&incomeCategories).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeCategoryGetFailed)
		return
	}
	output := IncomeCategoryGetsOutput{
		ResponseBase: command_models.Success,
		IncomeCategories: incomeCategories,
		Total: total,
	}
	ctx.JSON(200, output)
}

func IncomeCategoryDelete(ctx *gin.Context) {
	if err := db.DB.Where("id = ?", ctx.Param("income_category_id")).Delete(&models.IncomeCategory{}).Error; err != nil {
		ctx.JSON(200, resp_code.IncomeCategoryDeleteFailed)
		return
	}
	ctx.JSON(200, command_models.Success)
}
