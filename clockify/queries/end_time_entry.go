package queries

import (
	"bytes"
	"clk/clockify"
	"clk/db/models"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func EndTimeEntry(todo models.ToDo) (clockify.TimeEntry, error) {
	var timeEntry clockify.TimeEntry

	if todo.End == nil || todo.TimeEntryID == nil {
		return timeEntry, errors.New("Missing either .End or .TimeEntryID fields in todo")
	}

	requestJSON := map[string]string{
		"start":       todo.Start.UTC().Format(time.RFC3339),
		"description": todo.Description,
		"projectId":   todo.ProjectID,
		"end":         todo.End.UTC().Format(time.RFC3339),
	}
	byteBuffer := new(bytes.Buffer)
	json.NewEncoder(byteBuffer).Encode(requestJSON)

	fmt.Println("Putting to:", fmt.Sprint("workspaces/", todo.WorkspaceID, "/time-entries/", *todo.TimeEntryID))

	resp, err := clockify.Call("PUT", clockify.Api(fmt.Sprint("workspaces/", todo.WorkspaceID, "/time-entries/", *todo.TimeEntryID)), byteBuffer)
	if err != nil {
		return timeEntry, err
	}

	err = json.NewDecoder(resp.Body).Decode(&timeEntry)
	return timeEntry, err
}
