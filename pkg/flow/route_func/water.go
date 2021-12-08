package route_func

import (
	"github.com/gin-gonic/gin"
	"yang-backend/pkg/command/command_func"
	models2 "yang-backend/pkg/command/models"
	"yang-backend/pkg/command/resp_code"
	"yang-backend/pkg/db"
	"yang-backend/pkg/flow/models"
)

type CreateWaterInput struct {
	Water models.Water `json:"water"`
}

type CreateWaterOutput struct {
	models2.ResponseBase
	Data models.Water `json:"data"`
}

func CreateWater(ctx *gin.Context) {
	var input CreateWaterInput
	if err := ctx.BindJSON(&input); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&input.Water).Error; err != nil {
		ctx.JSON(200, resp_code.CreateWaterFailed)
		return
	}
	o := CreateWaterOutput{
		ResponseBase: models2.Success,
		Data:         input.Water,
	}
	ctx.JSON(200, o)
}

type CreateWaterCltInput struct {
	WaterClt models.WaterCollection `json:"water_clt"`
}

type CreateWaterCltOutput struct {
	models2.ResponseBase
	Data models.WaterCollection `json:"data"`
}

func CreateWaterClt(ctx *gin.Context) {
	var i CreateWaterCltInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&i.WaterClt).Error; err != nil {
		ctx.JSON(200, resp_code.CreateWaterCltFailed)
		return
	}
	o := CreateWaterCltOutput{
		ResponseBase: models2.Success,
		Data:         i.WaterClt,
	}
	ctx.JSON(200, o)
}

type CreateSortKlgInput struct {
	SortKnowledge models.SortKnowledge `json:"sort_knowledge"`
}
type CreateSortKlgOutput struct {
	models2.ResponseBase
	Data models.SortKnowledge `json:"data"`
}

func CreateSortKlg(ctx *gin.Context) {
	var i CreateSortKlgInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&i.SortKnowledge).Error; err != nil {
		ctx.JSON(200, resp_code.CreateSortKlgFailed)
		return
	}
	o := CreateSortKlgOutput{
		ResponseBase: models2.Success,
		Data:         i.SortKnowledge,
	}
	ctx.JSON(200, o)
}

type CreateTodoInput struct {
	Todo models.Todo `json:"todo"`
}
type CreateTodoOutput struct {
	models2.ResponseBase
	Data models.Todo `json:"data"`
}

func CreateTodo(ctx *gin.Context) {
	var i CreateTodoInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&i.Todo).Error; err != nil {
		ctx.JSON(200, resp_code.CreateTodoFailed)
		return
	}
	o := CreateTodoOutput{
		ResponseBase: models2.Success,
		Data:         i.Todo,
	}
	ctx.JSON(200, o)
}

type CreateDevTestInput struct {
	DevTest models.DevTest `json:"dev_test"`
}
type CreateDevTestOutput struct {
	models2.ResponseBase
	Data models.DevTest `json:"data"`
}

func CreateDevTest(ctx *gin.Context) {
	var i CreateDevTestInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Create(&i.DevTest).Error; err != nil {
		ctx.JSON(200, resp_code.CreateDevTestFailed)
		return
	}
	o := CreateDevTestOutput{
		ResponseBase: models2.Success,
		Data:         i.DevTest,
	}
	ctx.JSON(200, o)
}

type GetWatersInput struct {
	models2.GetListModel
}

type GetWatersOutput struct {
	models2.ResponseBase
	Data  []models.Water `json:"data"`
	Total int64          `json:"total"`
}

func GetWaters(ctx *gin.Context) {
	var i GetWatersInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(models.Water{}, i.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.GetWatersFailed)
		return
	}
	var waters []models.Water
	if err := tx.Find(&waters).Error; err != nil {
		ctx.JSON(200, resp_code.GetWatersFailed)
		return
	}
	o := GetWatersOutput{
		ResponseBase: models2.Success,
		Data:         waters,
		Total:        total,
	}
	ctx.JSON(200, o)
}

type GetWaterOutput struct {
	models2.ResponseBase
	Data models.Water `json:"data"`
}

func GetWater(ctx *gin.Context) {
	waterId := ctx.Param("water_id")
	var water models.Water
	if err := db.DB.Preload("Flow").Where("id = ?", waterId).First(&water).Error; err != nil {
		ctx.JSON(200, resp_code.GetWaterFailed)
		return
	}
	o := GetWaterOutput{
		ResponseBase: models2.Success,
		Data:         water,
	}
	ctx.JSON(200, o)
}

type GetSortKlgsInput struct {
	models2.GetListModel
	WaterID string `json:"water_id"`
}
type GetSortKlgsOutput struct {
	models2.ResponseBase
	Data  []models.SortKnowledge `json:"data"`
	Total int64                  `json:"total"`
}

func GetSortKlgs(ctx *gin.Context) {
	var i GetSortKlgsInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(&models.SortKnowledge{}, i.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.GetSortKlgsFailed)
		return
	}
	if i.WaterID != "" {
		tx.Where("water_id = ?", i.WaterID)
		tx.Count(&total)
	}
	var sortKlgs []models.SortKnowledge
	if err := tx.Find(&sortKlgs).Error; err != nil {
		ctx.JSON(200, resp_code.GetSortKlgsFailed)
		return
	}
	o := GetSortKlgsOutput{
	    ResponseBase: models2.Success,
	    Data: sortKlgs,
	    Total: total,
    }
    ctx.JSON(200, o)
}
