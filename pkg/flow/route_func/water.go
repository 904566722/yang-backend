package route_func

import (
	"errors"
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
	// 添加标签
	for index, _ := range i.WaterClt.WaterCltLabels {
		if err := getWaterCltLabelByName(i.WaterClt.WaterCltLabels[index].Name, &i.WaterClt.WaterCltLabels[index]); err != nil {
			ctx.JSON(200, resp_code.CreateWaterCltFailed)
			return
		}
	}

	if err := db.DB.Save(&i.WaterClt).Error; err != nil {
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

func getWaterCltLabelByName(name string, label *models.WaterCltLabel) error {
	var labels []models.WaterCltLabel
	if err := db.DB.Where("name = ?", name).Find(&labels).Error; err != nil {
		return errors.New("get water clt label failed")
	}
	if len(labels) > 0 {
		*label = labels[0]
	}
	return nil
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

type UpdateWaterInput struct {
	Water models.Water `json:"water"`
}
type UpdateWaterOutput struct {
	models2.ResponseBase
	Water models.Water `json:"water"`
}

func UpdateWater(ctx *gin.Context) {
	var waters []models.Water
	id := ctx.Param("water_id")
	if err := db.DB.Where("id = ?", id).Find(&waters).Error; err != nil {
		ctx.JSON(200, resp_code.UpdateWaterFailed)
		return
	}
	if len(waters) == 0 {
		ctx.JSON(200, resp_code.NotFoundResource)
		return
	}
	oldWater := waters[0]
	var i UpdateWaterInput
	i.Water = oldWater
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Save(&i.Water).Error; err != nil {
		ctx.JSON(200, resp_code.UpdateWaterFailed)
		return
	}
	o := UpdateWaterOutput{
		ResponseBase: models2.Success,
		Water:        i.Water,
	}
	ctx.JSON(200, o)
}

type UpdateTodoInput struct {
	Todo models.Todo `json:"todo"`
}
type UpdateTodoOutput struct {
	models2.ResponseBase
	Todo models.Todo `json:"todo"`
}

func UpdateTodo(ctx *gin.Context) {
	var todos []models.Todo
	id := ctx.Param("todo_id")
	if err := db.DB.Where("id = ?", id).Find(&todos).Error; err != nil {
		ctx.JSON(200, resp_code.UpdateTodoFailed)
		return
	}
	if len(todos) == 0 {
		ctx.JSON(200, resp_code.NotFoundResource)
		return
	}
	oldTodo := todos[0]
	var i UpdateTodoInput
	i.Todo = oldTodo
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Save(&i.Todo).Error; err != nil {
		ctx.JSON(200, resp_code.UpdateTodoFailed)
		return
	}
	o := UpdateTodoOutput{
		ResponseBase: models2.Success,
		Todo:         i.Todo,
	}
	ctx.JSON(200, o)
}

type UpdateWaterCltInput struct {
	WaterCollection models.WaterCollection `json:"water_collection"`
}
type UpdateWaterCltOutput struct {
	models2.ResponseBase
	WaterCollection models.WaterCollection `json:"water_collection"`
}

func UpdateWaterClt(ctx *gin.Context) {
	var waterClts []models.WaterCollection
	id := ctx.Param("water_clt_id")
	if err := db.DB.Where("id = ?", id).Find(&waterClts).Error; err != nil {
		ctx.JSON(200, resp_code.UpdateWaterFailed)
		return
	}
	if len(waterClts) == 0 {
		ctx.JSON(200, resp_code.NotFoundResource)
		return
	}
	oldWaterClt := waterClts[0]
	var i UpdateWaterCltInput
	i.WaterCollection = oldWaterClt
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	if err := db.DB.Save(&i.WaterCollection).Error; err != nil {
		ctx.JSON(200, resp_code.UpdateWaterCltFailed)
		return
	}
	o := UpdateWaterCltOutput{
		ResponseBase:    models2.Success,
		WaterCollection: i.WaterCollection,
	}
	ctx.JSON(200, o)
}

type DeleteWaterCltOutput struct {
	models2.ResponseBase
}

func DeleteWaterClt(ctx *gin.Context)  {
	id := ctx.Param("water_clt_id")
	if err := db.DB.Where("id = ?", id).Delete(&models.WaterCollection{}).Error; err != nil {
		ctx.JSON(200, resp_code.DeleteWaterCltFailed)
		return
	}
	o := DeleteWaterCltOutput{
		ResponseBase: models2.Success,
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
		Data:         sortKlgs,
		Total:        total,
	}
	ctx.JSON(200, o)
}

type GetTodosInput struct {
	models2.GetListModel
	WaterID string `json:"water_id"`
}
type GetTodosOutput struct {
	models2.ResponseBase
	Data  []models.Todo `json:"data"`
	Total int64         `json:"total"`
}

func GetTodos(ctx *gin.Context) {
	var i GetTodosInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(&models.Todo{}, i.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.GetTodosFailed)
		return
	}
	if i.WaterID != "" {
		tx.Where("water_id = ?", i.WaterID)
		tx.Count(&total)
	}
	var todos []models.Todo
	if err := tx.Find(&todos).Error; err != nil {
		ctx.JSON(200, resp_code.GetTodosFailed)
		return
	}
	o := GetTodosOutput{
		ResponseBase: models2.Success,
		Data:         todos,
		Total:        total,
	}
	ctx.JSON(200, o)
}

type GetWaterCltsInput struct {
	models2.GetListModel
	WaterId string `json:"water_id"`
}

type GetWaterCltsOutput struct {
	models2.ResponseBase
	Data  []models.WaterCollection `json:"data"`
	Total int64                    `json:"total"`
}

func GetWaterClts(ctx *gin.Context) {
	var i GetWaterCltsInput
	if err := ctx.BindJSON(&i); err != nil {
		inputError := models2.InputError(err)
		ctx.JSON(inputError.HttpCode(), inputError)
		return
	}
	tx, total, err := command_func.CommandGets(&models.WaterCollection{}, i.GetListModel)
	if err != nil {
		ctx.JSON(200, resp_code.GetWaterCltsFailed)
		return
	}
	if i.WaterId != "" {
		tx = tx.Where("water_id = ?", i.WaterId)
		tx.Count(&total)
	}
	// 分页
	pageIndex := i.PageIndex
	pageSize := i.PageSize
	if pageIndex != 0 && pageSize != 0 {
		tx = tx.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}
	var waterClts []models.WaterCollection
	if err := tx.Find(&waterClts).Error; err != nil {
		ctx.JSON(200, resp_code.GetWaterCltsFailed)
		return
	}

	o := GetWaterCltsOutput{
		ResponseBase: models2.Success,
		Data:         waterClts,
		Total:        total,
	}
	ctx.JSON(200, o)
}
