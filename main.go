package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")

	x := 10

	if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is greater than or equal to 10")
	}

	numbers := [5]int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	now := time.Now()
	fmt.Println("Date time: ", now)

	weekday := now.Weekday().String()
	fmt.Println("weekday: ", weekday)

	day := now.Day()
	fmt.Println("Day: ", day)

	month := now.Month()
	fmt.Println("Month: ", month)

	year := now.Year()
	fmt.Println("Year: ", year)

	switch weekday {
	case "Monday":
		fmt.Println("Monday Uff")
	case "Tuesday":
		fmt.Println("Tuesday Ok")
	case "Wednesday":
		fmt.Println("Wednesday tired already?")
	case "Thursday":
		fmt.Println("Thursday almost there ...")
	case "Friday", "Saturday":
		fmt.Println(weekday + " It's the weekeeeend!")
	default:
		fmt.Println(weekday + " enjoy while it lasts ...")
	}

	// Create a date in the past
	pastDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Calculate the difference between the two dates
	diff := now.Sub(pastDate)
	days := int(diff.Hours() / 24)
	fmt.Println("Days between now and January 1, 2023:", days)

}
