package models

import (
	"gorm.io/gorm"
	"time"
)

type CommandModel struct {
	ID        string         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Remark    string         `json:"remark"`
}

type GetListModel struct {
	PageIndex  int  `json:"page_index,omitempty"`
	PageSize   int  `json:"page_size,omitempty"`
	ShowDelete bool `json:"show_delete"`
}

type PageMent struct {
	PageIndex int `json:"page_index,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
}


