package models

import (
	"github.com/goravel/framework/database/orm"
)

type Project struct {
	orm.Model
	WorkspaceID uint       `gorm:"column:workspace_id" json:"workspace_id"`
	Name        string     `gorm:"column:name" json:"name"`
	Workspace   *Workspace `gorm:"foreignKey:WorkspaceID" json:"workspace,omitempty"`
	Tasks       []Task     `gorm:"foreignKey:ProjectID" json:"tasks,omitempty"`
}
