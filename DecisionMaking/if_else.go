package main

import "fmt"

func main() {
	var age int = 26
	if age > 18 {
		fmt.Println("Eligible for vote")
	} else {
		fmt.Println("Not eligible for vote")
	}
}
