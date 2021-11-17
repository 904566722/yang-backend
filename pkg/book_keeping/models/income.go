package models

import (
	"gorm.io/gorm"
	"yang-backend/pkg/utils"
)

// Income 收入
type Income struct {
	CommandModel
	Amount           float32 `json:"amount"`
	IncomeCategoryID string  `json:"income_category_id"`
	incomeCategory   IncomeCategory
}

func (i *Income) TableName() string {
	return "income"
}

func (i *Income) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = utils.GenerateId("income", 10)
	return
}

// IncomeCategory 收入类别
type IncomeCategory struct {
	CommandModel
	Name string `json:"name"`
}

func (ic *IncomeCategory) TableName() string {
	return "income_category"
}

func (ic *IncomeCategory) BeforeCreate(tx *gorm.DB) (err error) {
	ic.ID = utils.GenerateId("income-c", 10)
	return
}
