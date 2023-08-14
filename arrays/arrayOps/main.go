package main

import "fmt"

func main() {
	fmt.Println("Array Ops")

	ar := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ar)
	// NOTE: modifying an array element

	ar[0] = 100
	fmt.Println(ar)

	// NOTE: Ranging through an array

	for i, v := range ar {
		fmt.Printf("Item %d was found at index %d\n", v, i)
	}
	// NOTE: using a for loop instead of range
	fmt.Println("Using a loop")
	for i := 0; i < len(ar); i++ {
		fmt.Println("Index: ", i, "Value: ", ar[i])
	}
	// NOTE: Multidimenssional Arrays
	// multi is a Multidimenssional array
	multi := [2][3]int{
		{0, 1, 2}, {3, 4, 5},
	}
	fmt.Println(multi)
}
