package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Fullname   string      `gorm:"column:fullname" json:"fullname"`
	Username   string      `gorm:"column:username;unique" json:"username"`
	Email      string      `gorm:"column:email;unique" json:"email"`
	Password   string      `gorm:"column:password" json:"-"` // Hidden from JSON responses
	Workspaces []Workspace `gorm:"foreignKey:Owner" json:"workspaces,omitempty"`
}
