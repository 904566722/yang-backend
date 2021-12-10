package command_func

import (
	"gorm.io/gorm"
	"time"
	models2 "yang-backend/pkg/command/models"
	"yang-backend/pkg/config"
	"yang-backend/pkg/db"
)

// CommandGets 获取模型列表时通用的方法
func CommandGets(model interface{}, input models2.GetListModel) (tx *gorm.DB, total int64, err error) {
	// 是否显示软删除记录
	tx = db.DB.Model(model)
	if input.ShowDelete == true {
		tx = tx.Unscoped()
	}
	// 统计总数
	tx.Count(&total)
	// 分页
	pageIndex := input.PageIndex
	pageSize := input.PageSize
	if pageIndex != 0 && pageSize != 0 {
		tx = tx.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}
	// 预加载
	for _, association := range input.Associations {
		if err := tx.Preload(association.Name).Error; err != nil {
			continue
		}
		if association.BelongYear != 0 && association.BelongMon != 0 {
			beginAt, endAt := config.GetMonRange(association.BelongYear, association.BelongMon)
			tx.Preload(association.Name, "created_at >= ? and created_at <= ?", beginAt, endAt)
		} else if association.BeginAt != config.NilTime && association.EndAt != config.NilTime {
			tx.Preload(association.Name, "created_at >= ? and created_at <= ?", association.BeginAt, association.EndAt)
		} else {
			tx.Preload(association.Name)
		}
	}
	// 查询时间段
	nilTime := time.Time{}
	if input.BeginAt != nilTime && input.EndAt != nilTime {
		tx.Where("created_at >= ? and created_at <= ?", input.BeginAt, input.EndAt)
	}

	return tx, total, nil
}
