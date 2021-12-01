package models

import (
    "gorm.io/gorm"
    "yang-backend/pkg/command/models"
    "yang-backend/pkg/command/utils"
)

type Collection struct {
    models.CommandModel
    Content string `json:"content"`
}

func (c *Collection) TableName() string {
    return "collection"
}

func (c *Collection) BeforeCreate(tx *gorm.DB) (err error) {
    c.ID = utils.GenerateId("collection", 10)
    return
}