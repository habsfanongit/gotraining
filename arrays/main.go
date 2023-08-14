package main

import "fmt"

func main() {
	//NOTE: arrays in go are fixed and their size cannot be modified
	//slices on  the other hand are arrays with flexible size
	//to add a new item to a slice, use the append function
	//
	nums := [3]int{1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Printf("%v\n", nums)
	fmt.Printf("%#v\n", nums)

	// TODO:Todo here
	// FIX: fix here
	// WARNING:?
	// PERF:performance?
	// NOTE: Initialize a float array with zeros
	// HACK: This is a hack

	zeros := [2]float64{}
	fmt.Printf("%#v\n", zeros)
	// Adding Items to the zeros array
	zeros[0] = 1.2
	zeros[1] = 2.2

	fmt.Printf("%#v\n", zeros)

	// NOTE: ellipses array Operator

	ar := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ar)

	// NOTE: Initialize an array using multi lines like below must end with a comma
	// missing the last comma will result in an error

	ar1 := [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(ar1)
}
