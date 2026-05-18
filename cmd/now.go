package cmd

import (
	"prayer-cli/prayer"
	"prayer-cli/utils"

	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Show current and next prayer",
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

		current, next := prayer.GetCurrentAndNext(schedule)
		remaining := ""
		if next != nil{
			remaining = prayer.TimeRemaining(next)
		}

		utils.PrintNow(getLocationString(loc), current, next, remaining)
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)

	nowCmd.Flags().StringP("city", "c", "", "City name (e.g. 'New York')")
	nowCmd.Flags().StringP("country", "C", "", "Country name or code (e.g. 'US')")
}