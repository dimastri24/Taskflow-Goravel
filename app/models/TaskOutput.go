package models

import (
	"github.com/goravel/framework/database/orm"
)

type TaskOutput struct {
	orm.Model
	TaskID   uint    `gorm:"column:task_id" json:"task_id"`
	Type     int16   `gorm:"column:type" json:"type"`
	FileName *string `gorm:"column:file_name" json:"file_name"`
	FileExt  *string `gorm:"column:file_ext" json:"file_ext"`
	FileSize *int    `gorm:"column:file_size" json:"file_size"`
	Link     *string `gorm:"column:link" json:"link"`
	Task     *Task   `gorm:"foreignKey:TaskID" json:"task,omitempty"`
}
