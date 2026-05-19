package cmd

import (
	"prayer-cli/prayer"
	"prayer-cli/utils"

	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Show the next prayer and time remaining",
	Run: func(cmd *cobra.Command, args []string){
		city, _ := cmd.Flags().GetString("city")
		country, _ := cmd.Flags().GetString("country")

		loc, err := resolveLocation(city, country)
		if err != nil{
			exit(err)
		}

		schedule, err := fetchSchedule(loc)
		if err != nil{
			exit(err)
		}

		_, next := prayer.GetCurrentAndNext(schedule)
		remaining := ""
		if next != nil{
			remaining = prayer.TimeRemaining(next)
		}

		utils.PrintNext(getLocationString(loc), next, remaining)
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}