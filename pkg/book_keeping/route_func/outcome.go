package route_func

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yang-backend/pkg/book_keeping/models"
	"yang-backend/pkg/command/command_func"
	command_models "yang-backend/pkg/command/models"
	"yang-backend/pkg/command/resp_code"
	"yang-backend/pkg/command/utils"
	"yang-backend/pkg/config"
	"yang-backend/pkg/db"
)

type CreateOutcomeInput struct {
	Outcome   models.Outcome `json:"outcome"`
	Yesterday bool           `json:"yesterday"`
}

type CreateOutcomeResponse struct {
	command_models.ResponseBase
	Data models.Outcome `json:"data"`
}

func CreateOutcome(ctx *gin.Context) {
	var input CreateOutcomeInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if input.Yesterday {
		input.Outcome.CreatedAt = utils.GetYesterday()
	}
	if err := db.DB.Create(&input.Outcome).Error; err != nil {
		ctx.JSON(200, resp_code.OutcomeCreateFailed)
		return
	}
	resp := CreateOutcomeResponse{
		ResponseBase: command_models.Success,
		Data:         input.Outcome,
	}
	ctx.JSON(200, resp)
}

func DeleteOutcome(ctx *gin.Context) {
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

func UpdateOutcome(ctx *gin.Context) {
	outcomes, err := getOutcomeById(ctx.Param("outcome_id"))
	if err == nil && len(outcomes) == 0 {
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
	if err := db.DB.Save(&input.Outcome).Error; err != nil {
		ctx.JSON(200, resp_code.OutcomeUpdateFailed)
		return
	}
	o := UpdateOutcomeOutput{
		ResponseBase: command_models.Success,
		Outcome:      input.Outcome,
	}
	ctx.JSON(200, o)
}

type GetOutcomesInput struct {
	command_models.GetListModel
	OutcomeCategoryId string `json:"outcome_category_id"`
	BelongYear        int    `json:"belong_year"`
	BelongMon         int    `json:"belong_mon"`
	BelongDay         int    `json:"belong_day"`
}

type GetOutcomesOutput struct {
	command_models.ResponseBase
	Data  []models.Outcome `json:"data"`
	Total int64            `json:"total"`
}

func GetOutcomes(ctx *gin.Context) {
	var input GetOutcomesInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(&models.Outcome{}, input.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.OutcomeGetsFailed)
		return
	}
	var outcomes []models.Outcome
	if input.OutcomeCategoryId != "" {
		tx.Where("outcome_category_id = ?", input.OutcomeCategoryId)
	}
	if input.BelongYear != 0 && input.BelongMon != 0 && input.BelongDay != 0 {
		beginAt, endAt := config.GetDayRange(input.BelongYear, input.BelongMon, input.BelongDay)
		tx.Where("created_at >= ? and created_at < ?", beginAt, endAt)
	}
	tx.Count(&total)
	if err := tx.Find(&outcomes).Error; err != nil {
		ctx.JSON(200, resp_code.OutcomeGetsFailed)
		return
	}
	resp := GetOutcomesOutput{
		ResponseBase: command_models.Success,
		Data:         outcomes,
		Total:        total,
	}
	ctx.JSON(200, resp)
}

type StatisticMonEatInput struct {
	BelongYear        int    `json:"belong_year"`
	BelongMon         int    `json:"belong_mon"`
	OutcomeCategoryId string `json:"outcome_category_id"`
}

type StatisticMonEatOutput struct {
	command_models.ResponseBase
	Data MonEat `json:"data"`
}

type MonEat struct {
	Dates       []string  `json:"dates"`
	Values      []float32 `json:"values"`
	DayAvg      float32   `json:"day_avg"`      // 日均
	AmountTotal float32   `json:"amount_total"` // 总计
}

func StatisticMonEat(ctx *gin.Context) {
	var input StatisticMonEatInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := command_models.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	var outcomes []models.Outcome
	beginAt, endAt := config.GetMonRange(input.BelongYear, input.BelongMon)
	if err := db.DB.Where("created_at >= ? and created_at <= ?", beginAt, endAt).
		Where("outcome_category_id = ?", input.OutcomeCategoryId).
		Find(&outcomes).Error; err != nil {
		ctx.JSON(200, resp_code.StatisticMonEatFailed)
		return
	}
	var dayFlag []int = make([]int, config.MonDay[input.BelongMon])
	var dateStrs []string = make([]string, config.MonDay[input.BelongMon])
	var values []float32 = make([]float32, config.MonDay[input.BelongMon])
	var amountTotal float32
	var dayAvg float32
	for _, value := range outcomes {
		day := value.CreatedAt.Day()
		dateStr := fmt.Sprintf("%d-%d-%d", value.CreatedAt.Year(), value.CreatedAt.Month(), value.CreatedAt.Day())
		if day <= 10 {
			dateStrs[config.MonDay[input.BelongMon]-11+day] = dateStr
			values[config.MonDay[input.BelongMon]-11+day] = value.Amount
		} else {
			dateStrs[day-11] = dateStr
			values[day-11] = value.Amount
		}
		if value.CreatedAt.Day() == config.MonDay[input.BelongMon] {
			dayFlag[0] = 1
		} else {
			dayFlag[value.CreatedAt.Day()] = 1
		}
		amountTotal += value.Amount
	}
	dayAvg = amountTotal / float32(config.MonDay[input.BelongMon])
	for day, flag := range dayFlag {
		if flag != 1 {
			if day <= 10 {
				nextY, nextM := config.GetNextMon(input.BelongYear, input.BelongMon)
				value := ""
				if day == 0 {
					value = fmt.Sprintf("%d-%d-%d", input.BelongYear, input.BelongMon, config.MonDay[input.BelongMon])
				} else {
					value = fmt.Sprintf("%d-%d-%d", nextY, nextM, day)
				}
				dateStrs[config.MonDay[input.BelongMon]-11+day] = value
				values[config.MonDay[input.BelongMon]-11+day] = 0
			} else {
				dateStrs[day-11] = fmt.Sprintf("%d-%d-%d", input.BelongYear, input.BelongMon, day)
				values[day-11] = 0
			}
		}
	}
	var output = StatisticMonEatOutput{
		ResponseBase: command_models.Success,
		Data: MonEat{
			Dates:       dateStrs,
			Values:      values,
			AmountTotal: amountTotal,
			DayAvg:      dayAvg,
		},
	}
	ctx.JSON(200, output)
}
