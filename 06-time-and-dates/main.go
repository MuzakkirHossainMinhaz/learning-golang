package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Go Time and Dates Demo!")

	currentTime := time.Now()
	fmt.Printf("Current date and time: %s\n", currentTime.Format("Monday, January 2, 2006 at 3:04:05 PM MST"))

	createdAt := time.Date(2026, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Printf("Example date in UTC: %s\n", createdAt.Format("Monday, January 2, 2006 at 3:04 PM MST"))
}
