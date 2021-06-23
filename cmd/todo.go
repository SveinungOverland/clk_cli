package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	subcmd_todo "clk/cmd/subcmd/todo"
)

// taskCmd represents the task command
var TodoCmd = &cobra.Command{
	Use:     "todo",
	Aliases: []string{"t"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo called")
	},
}

func init() {
	rootCmd.AddCommand(TodoCmd)

	subcmd_todo.RegisterCurrent(TodoCmd)
	subcmd_todo.RegisterEnd(TodoCmd)
	subcmd_todo.RegisterNew(TodoCmd)
	subcmd_todo.RegisterStart(TodoCmd)
	subcmd_todo.RegisterList(TodoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
