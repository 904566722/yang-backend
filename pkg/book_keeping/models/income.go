package models

import (
	"gorm.io/gorm"
	"yang-backend/pkg/command/models"
	utils2 "yang-backend/pkg/command/utils"
)

// Income 收入
type Income struct {
	models.CommandModel
	Amount float32 `json:"amount"`
	// 属于什么时候的收入
	Year  int `json:"year"`
	Month int `json:"month"`

	IncomeCategoryID string         `json:"income_category_id" gorm:"size:256"`
	IncomeCategory   IncomeCategory `json:"income_category"`
}

func (i *Income) TableName() string {
	return "income"
}

func (i *Income) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = utils2.GenerateId("income", 10)
	return
}

// IncomeCategory 收入类别
type IncomeCategory struct {
	models.CommandModel
	Name   string `json:"name"`
	OpUnit string `json:"op_unit" gorm:"comment:操作单位，月：m， 日：d"`

	// has many
	Incomes []Income `json:"incomes" gorm:"foreignKey:IncomeCategoryID;references:ID"`
}

func (ic *IncomeCategory) TableName() string {
	return "income_category"
}

func (ic *IncomeCategory) BeforeCreate(tx *gorm.DB) (err error) {
	ic.ID = utils2.GenerateId("income-c", 10)
	return
}
