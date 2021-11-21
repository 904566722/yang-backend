package command_func

import (
	"gorm.io/gorm"
	models2 "yang-backend/pkg/command/models"
	"yang-backend/pkg/db"
)

// CommandGets 获取模型列表时通用的方法
func CommandGets(model interface{}, input models2.GetListModel) (tx *gorm.DB, total int64, err error) {
	// 是否显示软删除记录
	preload := db.DB.Model(model)
	if input.ShowDelete == true {
		preload = preload.Unscoped()
	}
	// 统计总数
	preload.Count(&total)
	// 分页
	pageIndex := input.PageIndex
	pageSize := input.PageSize
	if pageIndex != 0 && pageSize != 0 {
		preload = preload.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}
	return preload, total, nil
}
