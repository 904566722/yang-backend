package route_func

import (
    "errors"
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/models"
    command_models "yang-backend/pkg/command/models"
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
	}
	if err := db.DB.Create(&input.IncomeCategory).Error; err != nil {
		err := errors.New("create income category failed")
		ctx.JSON(-1, err)
	}
    resp := CreateIncomeCategoryResponse{ResponseBase: command_models.Success}
    resp.Data = input.IncomeCategory
    ctx.JSON(200, resp)
}
