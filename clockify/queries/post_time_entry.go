package queries

import (
	"bytes"
	"clk/clockify"
	"clk/db/models"
	"encoding/json"
	"fmt"
	"time"
)

func PostTimeEntry(todo models.ToDo) (clockify.TimeEntry, error) {
	var timeEntry clockify.TimeEntry

	start := time.Now().UTC().Format(time.RFC3339)

	requestJSON := map[string]string{
		"start":       start,
		"description": todo.Description,
		"projectId":   todo.ProjectID,
	}
	byteBuffer := new(bytes.Buffer)
	json.NewEncoder(byteBuffer).Encode(requestJSON)

	resp, err := clockify.Call("POST", clockify.Api(fmt.Sprint("workspaces/", todo.WorkspaceID, "/time-entries")), byteBuffer)
	if err != nil {
		return timeEntry, err
	}

	err = json.NewDecoder(resp.Body).Decode(&timeEntry)
	return timeEntry, err
}
