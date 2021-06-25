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
	"clk/util"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:     "select",
	Aliases: []string{"s"},
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

		todo, err := util.SelectToDo(showInactive)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Println("Using:", todo)
		viper.Set("active_todo", todo.ID)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("Could not save configuration:", err.Error())
		}
	},
}

func RegisterSelect(todo *cobra.Command) {
	todo.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
