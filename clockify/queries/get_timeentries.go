package queries

import (
	"clk/clockify"
	"encoding/json"
	"errors"
	"time"

	"github.com/spf13/viper"
)

func GetTimeEntries(startTime, endTime *time.Time) ([]clockify.TimeEntry, error) {
	activeWorkspace := viper.GetString("workspace_id")
	userId := viper.GetString("user_id")

	if activeWorkspace == "" {
		return nil, errors.New("Missing workspace setting, set it using `clk workspace *workspace name*` or by using flag `--workspace *workspace name*`")
	}
	if userId == "" {
		return nil, errors.New("Missing user id")
	}

	path := clockify.
		ApiBuilder().
		PathVar("workspaces", activeWorkspace).
		PathVar("user", userId).
		Path("time-entries")

	if startTime != nil {
		path.Query("start", startTime.UTC().Format(time.RFC3339))
	}
	if endTime != nil {
		path.Query("end", endTime.UTC().Format(time.RFC3339))
	}

	resp, err := clockify.Call(
		"GET",
		path.Build(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var timeEntries []clockify.TimeEntry
	err = json.NewDecoder(resp.Body).Decode(&timeEntries)
	if err != nil {
		return nil, err
	}

	return timeEntries, nil
}
