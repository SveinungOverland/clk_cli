package cmd

import (
	"clk/util"
	"fmt"

	"github.com/spf13/cobra"
)

// workspaceCmd represents the workspace command
var WorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspace called")
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		workspaces, err := util.GetWorkSpaces()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		output := make([]string, len(workspaces))
		for i, workspace := range workspaces {
			output[i] = workspace.Name
		}
		return output, cobra.ShellCompDirectiveNoFileComp
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
