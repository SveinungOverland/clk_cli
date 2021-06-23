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
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:     "end",
	Aliases: []string{"e"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentTodoID := viper.GetString("active_todo")
		if currentTodoID == "" {
			fmt.Println("No current todo, select one or create a new todo")
			return
		}
		var currentTodo models.ToDo
		result := db.Client.First(&currentTodo, currentTodoID)
		if result.Error != nil {
			fmt.Println("SQL error:", result.Error.Error())
			return
		}

		endTime := time.Now()

		currentTodo.End = &endTime

		timeEntry, err := queries.EndTimeEntry(currentTodo)
		if err != nil {
			fmt.Println("Error ending todo in clockify:", err.Error())
			return
		}

		if timeEntry.TimeInterval.End == nil {
			fmt.Println("Did not properly update time entry")
			fmt.Printf("%+v\n", timeEntry)
			return
		}

		result = db.Client.Save(&currentTodo)
		if result.Error != nil {
			fmt.Println("SQL error:", result.Error.Error())
		}
	},
}

func RegisterEnd(todo *cobra.Command) {
	todo.AddCommand(endCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// endCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// endCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
