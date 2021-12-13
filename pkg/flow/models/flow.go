package models

import (
	"gorm.io/gorm"
	"yang-backend/pkg/command/models"
	"yang-backend/pkg/command/utils"
)

type Flow struct {
	models.CommandModel
	Name string `json:"name"`
}

func (f *Flow) TableName() string {
	return "flow"
}

func (f *Flow) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = utils.GenerateId("flow", 10)
	return
}

type Water struct {
	models.CommandModel
	Name     string `json:"name"`
	OverView string `json:"over_view"` // 概览
	Progress int    `json:"progress" gorm:"default:0"`

	FlowID string `json:"flow_id" gorm:"size:256"`
	Flow   Flow   `json:"flow"`
}

func (w *Water) TableName() string {
	return "water"
}

func (w *Water) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = utils.GenerateId("water", 10)
	return
}

// WaterCollection 碎片信息
type WaterCollection struct {
	models.CommandModel
	Content string `json:"content"`

	WaterID        string `json:"water_id" gorm:"size:256"`
	Water          Water
	WaterCltLabels []WaterCltLabel `json:"water_clt_labels" gorm:"many2many:water_collection_labels"`
}

type WaterCltLabel struct {
	models.CommandModel
	Name string `json:"name"`

	//WaterCollectionID string `json:"water_collection_id" gorm:"size:256"`
	WaterCollections []WaterCollection `json:"water_collection" gorm:"many2many:water_collection_labels"`
}

func (wcl *WaterCltLabel) TableName() string {
	return "water_clt_label"
}

func (wcl *WaterCltLabel) BeforeCreate(tx *gorm.DB) (err error) {
	if wcl.ID == "" {
		wcl.ID = utils.GenerateId("water-clt-label", 10)
	}
	return
}

func (wc *WaterCollection) TableName() string {
	return "water_collection"
}

func (wc *WaterCollection) BeforeCreate(tx *gorm.DB) (err error) {
	if wc.ID == "" {
		wc.ID = utils.GenerateId("water-clt", 10)
	}
	return
}

// SortKnowledge 待整理知识点
type SortKnowledge struct {
	models.CommandModel
	Title string `json:"title"`

	WaterID string `json:"water_id" gorm:"size:256"`
	Water   Water
}

func (sk *SortKnowledge) TableName() string {
	return "sort_knowledge"
}

func (sk *SortKnowledge) BeforeCreate(tx *gorm.DB) (err error) {
	sk.ID = utils.GenerateId("sk", 10)
	return
}

// Todo 勿忘 todo
type Todo struct {
	models.CommandModel
	Title string `json:"title"`
	Done  bool   `json:"done"`

	WaterID string `json:"water_id" gorm:"size:256"`
	Water   Water
}

func (td *Todo) TableName() string {
	return "todo"
}

func (td *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	td.ID = utils.GenerateId("todo", 10)
	return
}

// DevTest 开发测试用例
type DevTest struct {
	models.CommandModel
	Title       string `json:"title"`
	ExpectedRst string `json:"expected"`   // 预期结果
	ActualRst   string `json:"actual_rst"` // 实际结果
	Conclusion  string `json:"conclusion"` // 结论

	WaterID string `json:"water_id" gorm:"size:256"`
	Water   Water
}

func (test *DevTest) TableName() string {
	return "dev_test"
}

func (test *DevTest) BeforeCreate(tx *gorm.DB) (err error) {
	test.ID = utils.GenerateId("test", 10)
	return
}
