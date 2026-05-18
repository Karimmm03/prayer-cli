package cmd

import (
	"prayer-cli/utils"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's full prayer schedule",
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

		utils.PrintSchedule(getLocationString(loc), schedule)
	},
}

func init(){
	rootCmd.AddCommand(todayCmd)

	todayCmd.Flags().StringP("city", "c", "", "City name (e.g. 'New York')")
	todayCmd.Flags().StringP("country", "C", "", "Country name or code (e.g. 'US')")
}