package location

import (
	"testing"
)

func TestGetFromIP(t *testing.T) {
	loc, err := GetFromIP()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if loc.City == "" {
		t.Error("expected a city name, got empty string")
	}
	t.Logf("Detected location: %s, %s", loc.City, loc.Country)
}

func TestGetFromCity(t *testing.T) {
	loc := GetFromCity("Giza")
	if loc.City != "Giza" {
		t.Errorf("expected 'Giza', got '%s'", loc.City)
	}
}
