package models

import (
	"github.com/goravel/framework/database/orm"
)

type Workspace struct {
	orm.Model
	Name     string    `gorm:"column:name" json:"name"`
	Owner    uint      `gorm:"column:owner" json:"owner"`
	User     *User     `gorm:"foreignKey:Owner" json:"user,omitempty"`
	Members  []User    `gorm:"many2many:workspace_members;" json:"members,omitempty"`
	Projects []Project `gorm:"foreignKey:WorkspaceID" json:"projects,omitempty"`
}
