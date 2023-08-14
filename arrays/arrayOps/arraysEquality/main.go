package main

import "fmt"

func main() {
	arr1 := [3]int{0, 1, 2}
	arr2 := [3]int{0, 1, 2}
	// prints true
	fmt.Printf("%v\n", arr1 == arr2)
	// the blow also prints true
	arr3 := arr2
	fmt.Printf("%v\n", arr3 == arr2)
	var arr5 [5]int
	var arr4 [5]int
	fmt.Println(arr4 == arr5)
}
