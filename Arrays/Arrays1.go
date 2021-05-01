package main

import "fmt"

func main() {
	//var days []string
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days[0]) //Monday

	weekdays := days[0:5]
	fmt.Println(weekdays) //[Monday Tuesday Wednesday Thursday Friday]
}
