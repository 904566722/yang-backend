package route_func

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/models"
    command_models "yang-backend/pkg/command/models"
    "yang-backend/pkg/command/resp_code"
    "yang-backend/pkg/db"
)

type CreateOutcomeInput struct {
    Outcome models.Outcome `json:"outcome"`
}

type CreateOutcomeResponse struct {
    command_models.ResponseBase
    Data models.Outcome `json:"data"`
}

func CreateOutcome(ctx *gin.Context)  {
    var input CreateOutcomeInput
    if err := ctx.BindJSON(&input); err != nil {
        inputError := command_models.InputError(err)
        ctx.JSON(inputError.HttpCode(), inputError)
        return
    }
    if err := db.DB.Create(&input.Outcome).Error; err != nil {
        ctx.JSON(200, resp_code.OutcomeCreateFailed)
        return
    }
    resp := CreateOutcomeResponse{
        ResponseBase: command_models.Success,
        Data: input.Outcome,
    }
    ctx.JSON(200, resp)
}

func DeleteOutcome(ctx *gin.Context)  {
    id := ctx.Param("outcome_id")
    if err := db.DB.Where("id = ?", id).Delete(&models.Outcome{}).Error; err != nil {
        ctx.JSON(200, resp_code.OutcomeCreateFailed)
        return
    }
    ctx.JSON(200, command_models.Success)
}

type UpdateOutComeInput struct {
    Outcome models.Outcome `json:"outcome"`
}

type UpdateOutcomeOutput struct {
    command_models.ResponseBase
    Outcome models.Outcome `json:"outcome"`
}

func getOutcomeById(id string) ([]models.Outcome, error) {
    var outcomes []models.Outcome
    err := db.DB.Where("id = ?", id).Find(&outcomes).Error
    return outcomes, err
}

func UpdateOutcome(ctx *gin.Context)  {
    outcomes, err := getOutcomeById(ctx.Param("outcome_id"))
    if err==nil && len(outcomes)==0 {
        resp_code.NotFoundResource.Message = fmt.Sprintf("未找到支出条目, outcome_id: %s", ctx.Param("outcome_id"))
        ctx.JSON(200, resp_code.NotFoundResource)
        return
    } else if err != nil {
        ctx.JSON(200, resp_code.OutcomeGetFailed)
        return
    }
    oldOutcome := outcomes[0]
    var input UpdateOutComeInput
    input.Outcome = oldOutcome
    if err := ctx.BindJSON(&input); err != nil {
        inputErr := command_models.InputError(err)
        ctx.JSON(inputErr.HttpCode(), inputErr)
        return
    }
    if err := db.DB.Save(&input.Outcome).Error; err !=nil {
        ctx.JSON(200, resp_code.OutcomeUpdateFailed)
        return
    }
    o := UpdateOutcomeOutput{
        ResponseBase: command_models.Success,
        Outcome: input.Outcome,
    }
    ctx.JSON(200, o)
}
