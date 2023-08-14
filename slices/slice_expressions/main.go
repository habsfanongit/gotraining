package main

import "fmt"

func main() {
	// NOTE: Slicing an array returns a slice NOT an array
	// Example
	arr := [5]int{10, 20, 30, 40, 50}
	// NOTE: [low index: high index] while the low index is inclusive, the high index is exclusive
	s := arr[1:3] // NOTE: the expression is exclusive as it only stops at index 3 and not include it.
	// NOTE: the below line returns
	// The value of s is [20 30] and s is of type []int
	fmt.Printf("The value of s is %v and s is of type %T\n", s, s)
	// NOTE: Examples of Low and high index Slicing
	// arr[1:] is equal to arr[1:len(arr)]
	// arr[:] is equal to arr[0:len(arr)] -All items
	// arr[:3] is equal arr[0:3]
	// NOTE: using append with expression
	arr2 := [5]int{10, 20, 30, 40, 50}
	s2 := append(arr2[:3], 100)
	// NOTE: that the above creates a new slice usin the expression s[:3] and add 100
	// which results in [10,20,30,100] since the rest were not included in the expression
	fmt.Println(s2)
}
