package cmd

import (
	"fmt"
	"os"

	"prayer-cli/api"
	"prayer-cli/location"
	"prayer-cli/prayer"
)

func getLocationString(loc *location.Location) string{
	if loc.Country != ""{
		return fmt.Sprintf("%s, %s", loc.City, loc.Country)
	}
	return loc.City
}

func resolveLocation (city, country string) (*location.Location, error){
	if (city != "" && country == "") || (city == "" && country != "") {
		return nil, fmt.Errorf("you must provide both --city and --country together")
	}

	if city != "" && country != "" {
		return &location.Location{City: city, Country: country}, nil
	}

	loc, err := location.GetFromIP()
	if err != nil{
		return nil, fmt.Errorf("could not detect location: %w", err)
	}
	return loc, nil
}

func fetchSchedule (loc *location.Location) (*prayer.Schedule, error){
	data, err := api.GetPrayerTimes(loc.City, loc.Country)
	if err != nil{
		return nil, fmt.Errorf("could not fetch prayer times: %w", err)
	}

	schedule, err := prayer.BuildSchedule(data)
	if err != nil{
		return nil, fmt.Errorf("could not build schedule: %w", err)
	}

	return schedule, nil
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "  Error: %v\n\n", err)
	os.Exit(1)
}