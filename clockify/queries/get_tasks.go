package queries

import (
	"clk/clockify"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func GetTasks() ([]clockify.Task, error) {
	activeWorkspace := viper.GetString("workspace_id")
	activeProject := viper.GetString("project_id")
	if activeWorkspace == "" {
		return nil, errors.New("Missing workspace setting, set it using `clk workspace *workspace name*` or by using flag `--workspace *workspace name*`")
	}
	if activeProject == "" {
		return nil, errors.New("Missing project setting, set it using `clk project set *project name*` or by using flag `--project *project name*`")
	}

	resp, err := clockify.Call("GET", clockify.Api(fmt.Sprint("workspaces/", activeWorkspace, "/projects/", activeProject, "/tasks")), nil)
	if err != nil {
		return nil, err
	}

	var projects []clockify.Task
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
