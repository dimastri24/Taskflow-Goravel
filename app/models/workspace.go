package models

import (
	"github.com/goravel/framework/database/orm"
)

type Workspace struct {
	orm.Model
	Name      string    `gorm:"column:name" json:"name"`
	Owner     uint      `gorm:"column:owner" json:"owner"`
	UserOwner *User     `gorm:"foreignKey:Owner" json:"user_owner,omitempty"`
	Members   []User    `gorm:"many2many:workspace_members;" json:"members,omitempty"`
	Projects  []Project `gorm:"foreignKey:WorkspaceID" json:"projects,omitempty"`
}

type WorkspaceWithCounts struct {
	Workspace
	ProjectsCount int64 `json:"projects_count"`
	MembersCount  int64 `json:"members_count"`
}
