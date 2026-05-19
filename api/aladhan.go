package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type PrayerTimings struct{
	Fajr    string `json:"Fajr"`
	Sunrise string `json:"Sunrise"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type PrayerData struct{
	Timings  PrayerTimings
	Timezone string
}

type apiResponse struct{
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct{
		Timings PrayerTimings `json:"timings"`
		Meta    struct {
			Timezone string `json:"timezone"`
		} `json:"meta"`
	} `json:"data"`
}

func GetPrayerTimes(city, country string) (*PrayerData, error){
	params := url.Values{}
	params.Set("city", city)
	params.Set("method", "5") // Muslim World League method

	if country != ""{
		params.Set("country", country)
	}

	fullURL := "https://api.aladhan.com/v1/timingsByCity?" + params.Encode()

	resp, err := http.Get(fullURL)
	if err != nil{
		return nil, fmt.Errorf("failed to reach AlAdhan API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("AlAdhan API returned status: %d", resp.StatusCode)
	}

	var result apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
		return nil, fmt.Errorf("failed to parse prayer times response: %w", err)
	}

	if result.Code != 200{
		return nil, fmt.Errorf("AlAdhan API error: %s", result.Status)
	}

	return &PrayerData{
		Timings:  result.Data.Timings,
		Timezone: result.Data.Meta.Timezone,
	}, nil
}