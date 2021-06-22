package cmd

import (
	"clk/clockify/queries"
	"clk/clockify/util"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// workspaceCmd represents the workspace command
var WorkspaceCmd = &cobra.Command{
	Use:   "workspace `workspace name`",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires a workspace name")
		} else if len(args) > 1 {
			return errors.New("Too many arguments")
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		workspaces, err := queries.GetWorkSpaces()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		output := make([]string, len(workspaces))
		for i, workspace := range workspaces {
			output[i] = workspace.Name
		}
		return output, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName := args[0]
		workspaceID, err := util.GetWorkspaceIDFromName(workspaceName)
		if err != nil {
			fmt.Println("Could not find workspace in your created workspaces, use one of:")
			workspaces, _ := queries.GetWorkSpaces()
			for _, workspace := range workspaces {
				fmt.Println(" -", workspace.Name)
			}
		}

		fmt.Println("Setting workspace to:", workspaceName)
		viper.Set("workspace_id", workspaceID)
		viper.Set("workspace_name", workspaceName)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("Could not save workspace to file:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(WorkspaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
