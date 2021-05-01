package main

import (
	"fmt"
)

func main() {
	weekdays := map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
	fmt.Println(weekdays["Sunday"]) // 7
	for index, day := range weekdays {
		fmt.Println("Key: ", day, "\t\t Value:", index)
	}
	// Key:  4                  Value: Thursday
	// Key:  5                  Value: Friday
	// Key:  6                  Value: Saturday
	// Key:  7                  Value: Sunday
	// Key:  1                  Value: Monday
	// Key:  2                  Value: Tuesday
	// Key:  3                  Value: Wednesday

}
