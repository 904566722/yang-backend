package route_func

import (
    "errors"
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/models"
    command_models "yang-backend/pkg/command/models"
    "yang-backend/pkg/db"
)

type CreateIncomeInput struct {
    Income models.Income `json:"income"`
}

type CreateIncomeResponse struct {
    command_models.ResponseBase
    Data models.Income `json:"data"`
}

func IncomeCreate(ctx *gin.Context)  {
    var input CreateIncomeInput
    if err := ctx.BindJSON(&input); err != nil {
        inputError := command_models.InputError(err)
        ctx.JSON(inputError.HttpCode(), inputError)
    }
    if err := db.DB.Create(&input.Income).Error; err != nil {
        err := errors.New("create income failed")
        ctx.JSON(-1, err)
    }
    resp := CreateIncomeResponse{ResponseBase: command_models.Success}
    resp.Data = input.Income
    ctx.JSON(200, resp)
}