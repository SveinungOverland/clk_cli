package util

import (
	"clk/clockify/queries"
	"errors"
)

func GetProjectIDFromName(projectName string) (projectID string, err error) {
	projects, err := queries.GetProjects()
	if err != nil {
		return
	}
	for _, project := range projects {
		if project.Name == projectName {
			projectID = project.ID
			return
		}
	}
	err = errors.New("No project of that name")
	return
}

func GetProjectNameFromID(projectID string) (projectName string, err error) {
	projects, err := queries.GetProjects()
	if err != nil {
		return
	}
	for _, project := range projects {
		if project.ID == projectID {
			projectName = project.Name
			return
		}
	}
	err = errors.New("No project with that ID")
	return
}
