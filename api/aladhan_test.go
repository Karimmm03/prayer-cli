package api

import (
	"testing"
)

func TestGetPrayerTimes(t *testing.T){
	timings, err := GetPrayerTimes("Cairo", "Egypt")
	if err != nil{
		t.Fatalf("expected no error, got: %v", err)
	}

	if timings.Fajr == ""{
		t.Error("expected Fajr time, got empty string")
	}
	if timings.Maghrib == ""{
		t.Error("expected Maghrib time, got empty string")
	}

	t.Logf("Fajr:    %s", timings.Fajr)
	t.Logf("Sunrise: %s", timings.Sunrise)
	t.Logf("Dhuhr:   %s", timings.Dhuhr)
	t.Logf("Asr:     %s", timings.Asr)
	t.Logf("Maghrib: %s", timings.Maghrib)
	t.Logf("Isha:    %s", timings.Isha)
}

func TestGetPrayerTimesManualCity(t *testing.T){
	timings, err := GetPrayerTimes("New York", "US")
	if err != nil{
		t.Fatalf("expected no error, got: %v", err)
	}

	if timings.Fajr == ""{
		t.Error("expected Fajr time, got empty string")
	}

	t.Logf("New York - Fajr: %s, Maghrib: %s", timings.Fajr, timings.Maghrib)
}