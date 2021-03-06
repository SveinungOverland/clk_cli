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
	"clk/db/queries"
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showInactive, err := cmd.Flags().GetBool("all")
		if err != nil {
			fmt.Println("Flag error:", err.Error())
			return
		}

		todos, err := queries.GetToDos(showInactive, 0, 0)
		if err != nil {
			fmt.Println("SQL error:", err.Error())
			return
		}

		if showInactive {
			fmt.Println("Todos:")
		} else {
			fmt.Println("Active todos:")
		}

		for _, todo := range todos {
			fmt.Println(todo)
		}
	},
}

func RegisterList(todo *cobra.Command) {
	todo.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
