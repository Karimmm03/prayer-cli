package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct{
	City    string `json:"city"`
	Country string `json:"country"`
}

func GetFromIP() (*Location, error){
	url := "https://ipapi.co/json/"

	resp, err := http.Get(url)
	if err != nil{
		return nil, fmt.Errorf("failed to reach location API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("location API returned status: %d", resp.StatusCode)
	}

	var loc Location
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil{
		return nil, fmt.Errorf("failed to parse location response: %w", err)
	}

	if loc.City == "" {
		return nil, fmt.Errorf("could not detect city from IP")
	}

	return &loc, nil
}