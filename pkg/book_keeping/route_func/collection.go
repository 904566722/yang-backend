package route_func

import (
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/models"
    command_models "yang-backend/pkg/command/models"
    "yang-backend/pkg/db"
)

type CreateCollectionInput struct {
    Collection models.Collection `json:"collection"`
}

type CreateCollectionOutput struct {
    command_models.ResponseBase
    Data models.Collection `json:"data"`
}

func CreateCollection(ctx *gin.Context)  {
    var input CreateCollectionInput
    if err := ctx.BindJSON(&input); err != nil {
        inputError := command_models.InputError(err)
        ctx.JSON(inputError.HttpCode(), inputError)
        return
    }
    if err := db.DB.Create(&input.Collection).Error; err != nil {
        ctx.JSON(200, "create collection failed")
        return
    }
    ctx.JSON(200, CreateCollectionOutput{
        ResponseBase: command_models.Success,
        Data: input.Collection,
    })
}
