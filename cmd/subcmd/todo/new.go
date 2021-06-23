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
	"clk/db"
	"clk/db/models"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName := viper.GetString("workspace_name")
		projectName := viper.GetString("project_name")
		if workspaceName == "" || projectName == "" {
			fmt.Println("Missing configuration, either workspace or project is missing, use `--workspace --project` flags or set it with commands")
			fmt.Println(workspaceName, ">", projectName)
			return
		}
		fmt.Println("Creating todo for", workspaceName, ">", projectName)
		var prompt promptui.Prompt
		prompt = promptui.Prompt{
			Label: "Description",
		}
		description, err := prompt.Run()
		if err != nil {
			fmt.Println("Invalid description, what did you do???")
			return
		}
		todo := models.ToDo{
			Description:   description,
			WorkspaceID:   viper.GetString("workspace_id"),
			ProjectID:     viper.GetString("project_id"),
			Active:        true,
			WorkspaceName: workspaceName,
			ProjectName:   projectName,
		}
		result := db.Client.Create(&todo)
		if result.Error != nil {
			fmt.Println("Error creating todo:", result.Error.Error())
			return
		}
		viper.Set("active_todo", todo.ID)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("Error writing active todo to file:", err.Error())
		}
	},
}

func RegisterNew(todo *cobra.Command) {
	todo.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
