package models

import (
	"database/sql"
	"time"
)

type CommandModel struct {
	ID string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Remark string `json:"remark"`
}
