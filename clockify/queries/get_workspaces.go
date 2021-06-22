package queries

import (
	"clk/clockify"
	"encoding/json"
)

func GetWorkSpaces() ([]clockify.Workspace, error) {
	resp, err := clockify.Call("GET", clockify.Api("workspaces"), nil)
	if err != nil {
		return nil, err
	}
	var workspaces []clockify.Workspace
	err = json.NewDecoder(resp.Body).Decode(&workspaces)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
