package prayer

import (
	"testing"
	"time"

	"prayer-cli/api"
)

func TestBuildSchedule(t *testing.T){
	data := &api.PrayerData{
		Timezone: "UTC",
		Timings: api.PrayerTimings{
			Fajr:    "04:21",
			Sunrise: "06:01",
			Dhuhr:   "12:51",
			Asr:     "16:28",
			Maghrib: "19:42",
			Isha:    "21:10",
		},
	}

	schedule, err := BuildSchedule(data)
	if err != nil{
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(schedule.Prayers) != 6{
		t.Errorf("expected 6 prayers, got %d", len(schedule.Prayers))
	}

	fajr := schedule.Prayers[0]
	if fajr.Name != "Fajr"{
		t.Errorf("expected first prayer to be Fajr, got %s", fajr.Name)
	}
	if fajr.Time.Hour() != 4 || fajr.Time.Minute() != 21{
		t.Errorf("expected Fajr at 04:21, got %02d:%02d", fajr.Time.Hour(), fajr.Time.Minute())
	}

	t.Logf("Schedule built successfully with %d prayers", len(schedule.Prayers))
}

func TestGetCurrentAndNext(t *testing.T){
	now := time.Now()

	schedule := &Schedule{
		Prayers: []Prayer{
			{Name: "Fajr", Time: now.Add(-5 * time.Hour)},
			{Name: "Sunrise", Time: now.Add(-4 * time.Hour)},
			{Name: "Dhuhr", Time: now.Add(-2 * time.Hour)},
			{Name: "Asr", Time: now.Add(-1 * time.Hour)},
			{Name: "Maghrib", Time: now.Add(1 * time.Hour)},
			{Name: "Isha", Time: now.Add(3 * time.Hour)},
		},
	}

	current, next := GetCurrentAndNext(schedule)

	if current == nil{
		t.Fatal("expected a current prayer, got nil")
	}
	if next == nil{
		t.Fatal("expected a next prayer, got nil")
	}

	if current.Name != "Asr"{
		t.Errorf("expected current to be Asr, got %s", current.Name)
	}
	if next.Name != "Maghrib"{
		t.Errorf("expected next to be Maghrib, got %s", next.Name)
	}

	t.Logf("Current: %s, Next: %s", current.Name, next.Name)
	t.Logf("Time remaining: %s", TimeRemaining(next))
}

func TestGetCurrentAndNext_BeforeFajr(t *testing.T){
	now := time.Now()

	schedule := &Schedule{
		Prayers: []Prayer{
			{Name: "Fajr", Time: now.Add(1 * time.Hour)},
			{Name: "Sunrise", Time: now.Add(2 * time.Hour)},
			{Name: "Dhuhr", Time: now.Add(4 * time.Hour)},
			{Name: "Asr", Time: now.Add(5 * time.Hour)},
			{Name: "Maghrib", Time: now.Add(7 * time.Hour)},
			{Name: "Isha", Time: now.Add(9 * time.Hour)},
		},
	}

	current, next := GetCurrentAndNext(schedule)

	if current == nil || next == nil{
		t.Fatal("expected current and next prayers, got nil")
	}

	if current.Name != "Isha"{
		t.Errorf("expected current to be Isha (yesterday), got %s", current.Name)
	}
	if next.Name != "Fajr"{
		t.Errorf("expected next to be Fajr (today), got %s", next.Name)
	}
	
	expectedIshaTime := schedule.Prayers[5].Time.AddDate(0, 0, -1)
	if !current.Time.Equal(expectedIshaTime) {
		t.Errorf("expected current time to be yesterday's Isha")
	}
}

func TestGetCurrentAndNext_AfterIsha(t *testing.T){
	now := time.Now()

	schedule := &Schedule{
		Prayers: []Prayer{
			{Name: "Fajr", Time: now.Add(-9 * time.Hour)},
			{Name: "Sunrise", Time: now.Add(-7 * time.Hour)},
			{Name: "Dhuhr", Time: now.Add(-5 * time.Hour)},
			{Name: "Asr", Time: now.Add(-4 * time.Hour)},
			{Name: "Maghrib", Time: now.Add(-2 * time.Hour)},
			{Name: "Isha", Time: now.Add(-1 * time.Hour)},
		},
	}

	current, next := GetCurrentAndNext(schedule)

	if current == nil || next == nil{
		t.Fatal("expected current and next prayers, got nil")
	}

	if current.Name != "Isha"{
		t.Errorf("expected current to be Isha (today), got %s", current.Name)
	}
	if next.Name != "Fajr"{
		t.Errorf("expected next to be Fajr (tomorrow), got %s", next.Name)
	}

	expectedFajrTime := schedule.Prayers[0].Time.AddDate(0, 0, 1)
	if !next.Time.Equal(expectedFajrTime) {
		t.Errorf("expected next time to be tomorrow's Fajr")
	}
}