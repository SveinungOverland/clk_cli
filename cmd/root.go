package cmd

import (
	"clk/clockify/util"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile       string // Config file
	api_key       string // Clockify api key
	workspaceName string // Clockify workspace
	projectName   string // Clockify project
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clk",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd.GenZshCompletionFile("./_clk")
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clk_config.yaml)")

	rootCmd.PersistentFlags().StringVar(&api_key, "api_key", "", "Clockify api key")
	rootCmd.PersistentFlags().StringVarP(&workspaceName, "workspace", "w", "", "Clockify workspace to use")
	rootCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Clockify project to use")

	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api_key"))
	viper.BindPFlag("workspace_name", rootCmd.PersistentFlags().Lookup("workspace"))
	viper.BindPFlag("project_name", rootCmd.PersistentFlags().Lookup("project"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		cobra.CheckErr(err)

		// Search config in home directory with name ".clk_cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".clk_config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
		fmt.Println(err.Error())
		return
	}

	if flag := rootCmd.PersistentFlags().Lookup("workspace"); flag.Value.String() != "" {
		// Workspace flag is set and the id needs to be found
		fmt.Println("Finding workspace id from flag:", flag.Value.String())
		workspaceID, err := util.GetWorkspaceIDFromName(viper.GetString("workspace_name"))
		if err != nil {
			fmt.Println("Could not find workspace id from --workspace flag:", err.Error())
		}
		viper.Set("workspace_id", workspaceID)
	}

	if flag := rootCmd.PersistentFlags().Lookup("project"); flag.Value.String() != "" {
		// Workspace flag is set and the id needs to be found
		fmt.Println("Finding project id from flag:", flag.Value.String())
		projectID, err := util.GetProjectIDFromName(viper.GetString("project_name"))
		if err != nil {
			fmt.Println("Could not find project id from --project flag:", err.Error())
		}
		viper.Set("project_id", projectID)
	}

	if viper.GetString("api_key") == "" {
		fmt.Print("Missing api key from clockify, enter it: ")
		var input string
		fmt.Scanln(&input)
		viper.Set("api_key", input)
		if util.IsApiKeyUseable() {
			fmt.Println("Saving config to:", fmt.Sprint(home, "/.clk_config.yaml"))
			if err := viper.WriteConfigAs(fmt.Sprint(home, "/.clk_config.yaml")); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("Got here")
		} else {
			fmt.Println("Unusable api key, try another")
			os.Exit(1)
		}
	}
}
