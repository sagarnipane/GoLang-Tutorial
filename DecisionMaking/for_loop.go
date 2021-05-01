package main

import "fmt"

func main() {
	//Method 1
	var arr [5]int
	arr[0] = 2
	arr[1] = 2

	for i := range arr {
		fmt.Println(arr[i])
	}
	// 2 2 0 0 0

	arr2 := [4]int{1, 2, 3}
	for i := range arr2 {
		fmt.Println(arr2[i])
	}
	// 1 2 3 0
	// fmt.Println("\n Accessing array element out of index")
	// for k := 0; k < 5; k++ {
	// 	fmt.Print(arr2[k])
	// }
	// 	Accessing array element out of index
	// 1230panic: runtime error: index out of range [4] with length 4

	fmt.Println("\n Looping Over String")
	name := "sagar"
	for k := 0; k < 5; k++ {
		fmt.Println("%c", name[k]) //115 97 103 97 114
		// Println will print ASCII codes of a character/symbol
	}
	for k := 0; k < 5; k++ {
		fmt.Printf("%c", name[k]) //sagar
		//Because go-lang uses C language libraries it will print charaters
	}
	for k := 0; k < 5; k++ {
		fmt.Print(name[k]) // 1159710397114
		// Println will print ASCII codes of a character/symbol
	}
}
