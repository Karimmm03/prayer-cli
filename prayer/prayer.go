package prayer

import (
	"fmt"
	"time"

	"prayer-cli/api"
)

type Prayer struct{
	Name string
	Time time.Time
}

type Schedule struct{
	Prayers []Prayer
}

func BuildSchedule (timings *api.PrayerTimings)(*Schedule, error){
	rawTimes := []struct{
		name string
		raw string
	}{
		{"Fajr", timings.Fajr},
		{"Sunrise", timings.Sunrise},
		{"Dhuhr", timings.Dhuhr},
		{"Asr", timings.Asr},
		{"Maghrib", timings.Maghrib},
		{"Isha", timings.Isha},
	}
	var prayers []Prayer
	for _, p := range rawTimes{
		t, err := parseTime(p.raw)
		if err != nil{
			return nil, fmt.Errorf("failed to parse %s time: %w", p.name, err)
		}
		prayers = append(prayers, Prayer{Name: p.name, Time: t})
	}

	return &Schedule{Prayers: prayers}, nil
}

func parseTime (raw string)(time.Time, error){
	now := time.Now()
	if len(raw) > 5{
		raw = raw[:5] // sometimes api returns "HH:MM (EET)" with some info so we need only "HH:MM"
	}
	t, err := time.Parse("15:04", raw) // "15:04" is Go's ref time format for HH:MM
	if err != nil{
		return time.Time{}, err
	}

	return time.Date(
		now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location(),
	), nil
}

func GetCurrentAndNext (schedule *Schedule)(current, next *Prayer){
	now := time.Now()

	for i, p := range schedule.Prayers{
		if now.Before(p.Time){
			next = &schedule.Prayers[i]
			if i > 0{
				current = &schedule.Prayers[i-1]
			}
			return current, next
		}
	}

	last := &schedule.Prayers[len(schedule.Prayers) - 1]
	return last, nil
}

func TimeRemaining (next *Prayer) string{
	remaining := time.Until(next.Time)
	hours := int(remaining.Hours())
	minutes := int(remaining.Minutes()) % 60

	return fmt.Sprintf("%02dh %02dm", hours, minutes)
}