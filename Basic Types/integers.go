package main

import "fmt"

func main() {
	var myInt8 int8
	var myInt16 int16
	var myInt32 int32
	var myInt64 int64
	var myInt int
	myInt = 1
	myInt8 = 8
	myInt16 = 16
	myInt32 = 32
	myInt64 = 64
	fmt.Println(int8(myInt) + int8(myInt8))
	fmt.Println("Sagar " + string(rune(10)))
	fmt.Println(myInt16)
	fmt.Println(myInt32)
	fmt.Println(myInt64)

	myint2 := 10
	fmt.Println(myint2)
}
