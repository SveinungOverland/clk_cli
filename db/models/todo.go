package models

import (
	"fmt"
	"time"
)

type ToDo struct {
	ID            uint `gorm:"primarykey"`
	TimeEntryID   *string
	Start         *time.Time
	End           *time.Time
	Description   string
	WorkspaceID   string
	WorkspaceName string
	ProjectID     string
	ProjectName   string
	Active        bool      `gorm:"index:active_index,priority:1,sort:desc"`
	CreatedAt     time.Time `gorm:"index:active_index,priority:2,sort:desc"`
}

func (todo ToDo) String() string {
	var active string
	if todo.TimeEntryID != nil && todo.Start != nil && todo.End == nil {
		active = time.Since(*todo.Start).String()
	} else if todo.Active {
		active = "active"
	} else {
		active = "inactive"
	}
	return fmt.Sprintf("[%s] [%s > %s] %s", active, todo.WorkspaceName, todo.ProjectName, todo.Description)
}
