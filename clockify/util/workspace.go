package util

import (
	"clk/clockify/queries"
	"errors"
)

func GetWorkspaceIDFromName(workspaceName string) (workspaceID string, err error) {
	workspaces, err := queries.GetWorkSpaces()
	if err != nil {
		return
	}
	for _, workspace := range workspaces {
		if workspace.Name == workspaceName {
			workspaceID = workspace.ID
			return
		}
	}
	err = errors.New("No workspace of that name")
	return
}

func GetWorkspaceNameFromID(workspaceID string) (workspaceName string, err error) {
	workspaces, err := queries.GetWorkSpaces()
	if err != nil {
		return
	}
	for _, workspace := range workspaces {
		if workspace.ID == workspaceID {
			workspaceName = workspace.Name
			return
		}
	}
	err = errors.New("No workspace with that ID")
	return
}
