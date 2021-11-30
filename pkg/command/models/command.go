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
	PageIndex    int           `json:"page_index,omitempty"`
	PageSize     int           `json:"page_size,omitempty"`
	ShowDelete   bool          `json:"show_delete"`
	Associations []Association `json:"associations"`
	BeginAt      time.Time     `json:"begin_at"`
	EndAt        time.Time     `json:"end_at"`
}


// Association 加载关联关系
type Association struct {
	Name       string    `json:"name"`
	BelongYear int       `json:"belong_year"`   // 可以传入属于的 年份、月份
	BelongMon  int       `json:"belong_mon"`
	BeginAt    time.Time `json:"begin_at"`      // 也可以传入 开始时间、结束时间
	EndAt      time.Time `json:"end_at"`
}

type PageMent struct {
	PageIndex int `json:"page_index,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
}
