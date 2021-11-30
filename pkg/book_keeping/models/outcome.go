package models

import (
	"gorm.io/gorm"
	"yang-backend/pkg/command/models"
	utils2 "yang-backend/pkg/command/utils"
)

type Outcome struct {
	models.CommandModel
	Amount float32 `json:"amount"`
	Year   int     `json:"year"`
	Month  int     `json:"month"`

	OutcomeCategoryId string `json:"outcome_category_id" gorm:"size:256"`
	OutcomeCategory   OutcomeCategory
}

func (o *Outcome) TableName() string {
	return "outcome"
}

func (o *Outcome) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = utils2.GenerateId("outcome", 10)
	return
}

type OutcomeCategory struct {
	models.CommandModel
	Name string `json:"name"`
	OpUnit string `json:"op_unit" gorm:"comment:操作单位，月：m， 日：d"`

	Outcomes []Outcome `json:"outcomes" gorm:"foreignKey:OutcomeCategoryId;references:ID"`
}

func (oc *OutcomeCategory) TableName() string {
	return "outcome_category"
}

func (oc *OutcomeCategory) BeforeCreate(tx *gorm.DB) (err error) {
	oc.ID = utils2.GenerateId("outcome-c", 10)
	return
}
