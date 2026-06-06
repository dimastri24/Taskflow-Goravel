package models

import (
	"github.com/goravel/framework/database/orm"
)

type TaskComment struct {
	orm.Model
	UserID          uint          `gorm:"column:user_id" json:"user_id"`
	TaskID          uint          `gorm:"column:task_id" json:"task_id"`
	Message         string        `gorm:"column:message" json:"message"`
	ParentCommentID *uint         `gorm:"column:parent_comment_id" json:"parent_comment_id"`
	User            *User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Task            *Task         `gorm:"foreignKey:TaskID" json:"task,omitempty"`
	ParentComment   *TaskComment  `gorm:"foreignKey:ParentCommentID" json:"parent_comment,omitempty"`
	Replies         []TaskComment `gorm:"foreignKey:ParentCommentID" json:"replies,omitempty"`
}
