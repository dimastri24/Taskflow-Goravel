package models

import (
	"time"
)

type WorkspaceMember struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"column:user_id;uniqueIndex:idx_user_workspace" json:"user_id"`
	WorkspaceID uint      `gorm:"column:workspace_id;uniqueIndex:idx_user_workspace" json:"workspace_id"`
	JoinedAt    time.Time `gorm:"column:joined_at;type:timestamp" json:"joined_at"`
}
