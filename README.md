# Prayer CLI

A CLI tool that fetches Islamic prayer times based on your location — either auto-detected from your IP or manually specified — and displays the current prayer, next prayer, and full daily schedule directly in your terminal.

## Features
- **Auto-Detection**: Instantly detect your current location via IP.
- **Timezone Aware**: Accurately fetches and parses times for any specific timezone around the world.
- **Cross-day Tracking**: Seamlessly handles time boundaries (e.g., calculates tomorrow's Fajr if Isha has already passed).

---

## Installation

Ensure you have [Go installed](https://go.dev/doc/install) on your machine.

Clone the repository and build the binary:

```bash
git clone https://github.com/Karimmm03/prayer-cli.git
cd prayer-cli
go build -o pray.exe
```
*(Note: If you are on macOS or Linux, run `go build -o pray` instead and use `./pray` for the commands below).*

Move the binary to your PATH so you can run `pray` from anywhere:
 
```bash
# Linux / macOS
mv pray /usr/local/bin/
 
# Windows — move pray.exe to a folder that's in your PATH
```

---

## Usage

If you run a command without any flags, it will attempt to **auto-detect** your location based on your public IP address.

### Auto-detect your location
 
```bash
./pray.exe now        # current and next prayer
./pray.exe next       # next prayer and time remaining
./pray.exe today      # full prayer schedule for today
```
 
### Specify a location manually
 
```bash
./pray.exe now   --city "New York"    --country "US"
./pray.exe next  --city "London"      --country "GB"
./pray.exe today --city "Kuala Lumpur" --country "MY"
```

---

## Project Structure
 
```
prayer-cli/
├── main.go               # entry point
├── cmd/
│   ├── root.go           # base CLI command and persistent flags
│   ├── shared.go         # shared logic between commands
│   ├── now.go            # pray now
│   ├── next.go           # pray next
│   └── today.go          # pray today
├── api/
│   ├── aladhan.go        # AlAdhan API integration
│   └── aladhan_test.go   # API tests
├── location/
│   ├── location.go       # IP geolocation and manual city input
│   └── location_test.go  # location tests
├── prayer/
│   ├── prayer.go         # current/next logic and time parsing
│   └── prayer_test.go    # time and schedule tests
└── utils/
    └── format.go         # terminal output formatting
```

## APIs Used
 
- [AlAdhan API](https://aladhan.com/prayer-times-api) — prayer times by city
- [ipapi.co](https://ipapi.co) — IP geolocation (1000 free requests/day)

## Technical Details
This project is built in Go using the [Cobra CLI framework](https://github.com/spf13/cobra) and relies on the [AlAdhan API](https://aladhan.com/prayer-times-api) (Method 5: Muslim World League) for scheduling. 
