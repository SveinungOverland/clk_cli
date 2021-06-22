package queries

import (
	"clk/clockify"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func GetProjects() ([]clockify.Project, error) {
	activeWorkspace := viper.GetString("workspace_id")
	if activeWorkspace == "" {
		return nil, errors.New("Missing workspace setting, set it using `clk workspace *workspace name*` or by using flag `--workspace *workspace name*`")
	}
	resp, err := clockify.Call("GET", clockify.Api(fmt.Sprint("workspaces/", activeWorkspace, "/projects")), nil)
	if err != nil {
		return nil, err
	}
	var projects []clockify.Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
