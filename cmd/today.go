package cmd

import (
	"clk/clockify/queries"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var TodayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day := time.Now().Date()
		today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
		entries, err := queries.GetTimeEntries(&today, nil)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		for i := len(entries) - 1; i >= 0; i-- {
			place := len(entries) - i
			fmt.Println(place, " - ", entries[i])
		}

		var sumTime float64
		for _, entry := range entries {
			duration, err := time.ParseDuration(
				strings.ToLower(
					strings.TrimPrefix(*entry.TimeInterval.Duration, "PT"),
				),
			)
			if err != nil {
				fmt.Println("Error:", err.Error())
				return
			}
			sumTime += duration.Seconds()
		}

		if sumDuration, err := time.ParseDuration(fmt.Sprint(sumTime, "s")); err == nil {
			fmt.Println("Duration tracked today:", sumDuration)
		} else {
			fmt.Println("Error:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(TodayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
