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

func BuildSchedule (data *api.PrayerData)(*Schedule, error){
	rawTimes := []struct{
		name string
		raw string
	}{
		{"Fajr", data.Timings.Fajr},
		{"Sunrise", data.Timings.Sunrise},
		{"Dhuhr", data.Timings.Dhuhr},
		{"Asr", data.Timings.Asr},
		{"Maghrib", data.Timings.Maghrib},
		{"Isha", data.Timings.Isha},
	}

	var loc *time.Location
	if data.Timezone != ""{
		parsedLoc, err := time.LoadLocation(data.Timezone)
		if err == nil{
			loc = parsedLoc
		}
	}
	if loc == nil{
		loc = time.Local
	}

	var prayers []Prayer
	for _, p := range rawTimes{
		t, err := parseTime(p.raw, loc)
		if err != nil{
			return nil, fmt.Errorf("failed to parse %s time: %w", p.name, err)
		}
		prayers = append(prayers, Prayer{Name: p.name, Time: t})
	}

	return &Schedule{Prayers: prayers}, nil
}

func parseTime (raw string, loc *time.Location)(time.Time, error){
	if len(raw) > 5{
		raw = raw[:5] // sometimes api returns "HH:MM (EET)" with some info so we need only "HH:MM"
	}
	t, err := time.Parse("15:04", raw) // "15:04" is Go's ref time format for HH:MM
	if err != nil{
		return time.Time{}, err
	}
	
	now := time.Now().In(loc)
	return time.Date(
		now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, loc,
	), nil
}

func GetCurrentAndNext (schedule *Schedule)(current, next *Prayer){
	if len(schedule.Prayers) == 0{
		return nil, nil
	}
	
	loc := schedule.Prayers[0].Time.Location()
	now := time.Now().In(loc)

	for i, p := range schedule.Prayers{
		if now.Before(p.Time){
			next = &schedule.Prayers[i]
			if i > 0{
				current = &schedule.Prayers[i-1]
			} else {
				isha := schedule.Prayers[len(schedule.Prayers)-1]
				current = &Prayer{
					Name: isha.Name,
					Time: isha.Time.AddDate(0, 0, -1),
				}
			}
			return current, next
		}
	}

	last := &schedule.Prayers[len(schedule.Prayers) - 1]
	first := schedule.Prayers[0]
	
	nextDayFajr := &Prayer{
		Name: first.Name,
		Time: first.Time.AddDate(0, 0, 1),
	}

	return last, nextDayFajr
}

func TimeRemaining (next *Prayer) string{
	remaining := time.Until(next.Time)
	hours := int(remaining.Hours())
	minutes := int(remaining.Minutes()) % 60

	return fmt.Sprintf("%02dh %02dm", hours, minutes)
}