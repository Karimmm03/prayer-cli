package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "pray",
	Short: "A CLI tool that shows Islamic prayer times",
	Long: "pray - Get Islamic prayer times based on your location, right in your terminal.",
	
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Use 'pray now', 'pray next', or 'pray today'")
	},
}
func init() {
	rootCmd.PersistentFlags().StringP("city", "c", "", "City name (e.g. 'New York')")
	rootCmd.PersistentFlags().StringP("country", "C", "", "Country name or code (e.g. 'US')")
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}