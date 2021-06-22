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
package project

import (
	"clk/clockify/queries"
	"clk/clockify/util"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set PROJECT_NAME",
	Short: "Sets default project to given project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires a project name")
		} else if len(args) > 1 {
			return errors.New("Too many arguments")
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		projects, err := queries.GetProjects()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		output := make([]string, len(projects))
		for i, project := range projects {
			output[i] = project.Name
		}
		return output, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("workspace_name") == "" {
			fmt.Println("A workspace needs to be set, run `clk workspace `workspace name``")
			return
		}
		projectName := args[0]

		projectID, err := util.GetProjectIDFromName(projectName)
		if err != nil {
			fmt.Printf("Could not find project in workspace %s, use one of these or change workspace:\n", viper.GetString("workspace_name"))
			projects, _ := queries.GetProjects()
			for _, project := range projects {
				fmt.Println(" -", project.Name)
			}
		}

		fmt.Println("Setting workspace to:", projectName)
		viper.Set("project_id", projectID)
		viper.Set("project_name", projectName)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("Could not save project to file:", err.Error())
		}
	},
}

func RegisterSet(project *cobra.Command) {
	project.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
