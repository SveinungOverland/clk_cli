/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package todo

import (
	"clk/clockify/queries"
	"clk/db"
	"clk/db/models"
	"clk/util"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runningTodo := viper.GetUint("running_todo")
		if runningTodo != 0 {
			fmt.Println("A todo is already running, you must end it, setting it to selected todo for convenience")
			viper.Set("active_todo", runningTodo)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println("Error updating config:", err.Error())
				return
			}
		}
		selectTodo, err := cmd.Flags().GetBool("select")
		if err != nil {
			fmt.Println("Flag error:", err.Error())
			return
		}
		showInactive, err := cmd.Flags().GetBool("all")
		if err != nil {
			fmt.Println("Flag error:", err.Error())
			return
		}

		var currentTodoID uint
		var currentTodo models.ToDo

		// Use select instead of currently selected todo
		if selectTodo {
			currentTodo, err = util.SelectToDo(showInactive)
			if err != nil {
				fmt.Println("Error:", err.Error())
				return
			}
			currentTodoID = currentTodo.ID
		} else {
			currentTodoID = viper.GetUint("active_todo")
			if currentTodoID == 0 {
				fmt.Println("No current todo, select one or create a new todo")
				return
			}
			result := db.Client.First(&currentTodo, currentTodoID)
			if result.Error != nil {
				fmt.Println("SQL error:", result.Error.Error())
				return
			}
		}

		timeEntry, err := queries.PostTimeEntry(currentTodo)
		if err != nil {
			fmt.Println("Error when creating time entry:", err.Error())
			return
		}

		startTime, err := time.Parse(time.RFC3339, timeEntry.TimeInterval.Start)
		if err != nil {
			fmt.Println("Could not parse start time from clockify:", err.Error())
			return
		}

		currentTodo.TimeEntryID = &timeEntry.ID
		currentTodo.Start = &startTime
		currentTodo.End = nil
		currentTodo.Active = true
		result := db.Client.Save(&currentTodo)
		if result.Error != nil {
			fmt.Println("SQL error:", result.Error.Error())
		}
		viper.Set("running_todo", currentTodo.ID)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("Error updating config file:", err.Error())
			return
		}
		fmt.Println("Successfully started:", currentTodo)
	},
}

func RegisterStart(todo *cobra.Command) {
	todo.AddCommand(startCmd)

	startCmd.Flags().BoolP("select", "s", false, "Select which todo to start, instead of using currently selected one")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
