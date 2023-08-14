package main

import "fmt"

func main() {
	// Slices cannot be compared using the == operator.
	// the only way to compare a slice is by looping through
	// their elements

	a, b := []int{10, 2, 3}, []int{1, 2, 3}
	fmt.Println("a is ", a)
	a = nil
	areEqual := compare_slices(a, b)
	fmt.Println("The value of areEqual is ", areEqual)
}

func compare_slices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
