package models

import "time"

type ToDo struct {
	ID            uint `gorm:"primarykey"`
	TimeEntryID   *string
	Description   string
	WorkspaceID   string
	WorkspaceName string
	ProjectID     string
	ProjectName   string
	Active        bool      `gorm:"index:active_index,priority:1,sort:desc"`
	CreatedAt     time.Time `gorm:"index:active_index,priority:2,sort:desc"`
}
