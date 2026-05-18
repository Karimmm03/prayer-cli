package utils

import (
	"fmt"
	"prayer-cli/prayer"
)

func PrintSchedule(location string, schedule *prayer.Schedule){
	fmt.Println()
	fmt.Printf("  Location : %s\n", location)
	fmt.Println("  -------------------------")
	fmt.Printf("  %-10s %s\n", "Prayer", "Time")
	fmt.Println("  -------------------------")
	for _, p := range schedule.Prayers{
		fmt.Printf("  %-10s %s\n", p.Name, p.Time.Format("15:04"))
	}
	fmt.Println()
}

func PrintNow(location string, current, next *prayer.Prayer, remaining string){
	fmt.Println()
	fmt.Printf("  Location      : %s\n", location)
	fmt.Println("  -------------------------")
	if current != nil{
		fmt.Printf("  Current Prayer: %s\n", current.Name)
	} else {
		fmt.Println("  Current Prayer: --")
	}
	if next != nil{
		fmt.Printf("  Next Prayer   : %s at %s\n", next.Name, next.Time.Format("15:04"))
		fmt.Printf("  Time Remaining: %s\n", remaining)
	} else {
		fmt.Println("  Next Prayer   : No more prayers today")
	}
	fmt.Println()
}

func PrintNext(location string, next *prayer.Prayer, remaining string){
	fmt.Println()
	fmt.Printf("  Location      : %s\n", location)
	fmt.Println("  -------------------------")
	if next != nil{
		fmt.Printf("  Next Prayer   : %s\n", next.Name)
		fmt.Printf("  Time          : %s\n", next.Time.Format("15:04"))
		fmt.Printf("  Time Remaining: %s\n", remaining)
	} else {
		fmt.Println("  Next Prayer   : No more prayers today")
	}
	fmt.Println()
}