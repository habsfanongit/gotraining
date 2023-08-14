package main

import "fmt"

func main() {
	// NOTE: The append fucntion returns a new slice. It does not modify
	// existing ones.
	nums := []int{1, 2, 3, 4}
	// NOTE: The append function is variadic function
	nums = append(nums, 5, 6, 7, 8, 9, 10)
	fmt.Println(nums)
	// NOTE: copying Slices

	c, d := []int{1, 2, 3}, []int{4, 5, 6}
	// NOTE: Copying one slice to another using the ellipse operator
	c = append(c, d...)
	fmt.Println(c)

	// NOTE: Copying a slice using the copy function
	// copy is used to copy a source slice into a destination slice
	// and returns the number of elements copied
	src := []int{10, 20, 30}
	dest := make([]int, len(src))
	all := copy(dest, src)
	fmt.Printf("src is %v dest is %v and all is %v\n", src, dest, all)

	// NOTE: Using copy if the source and destination size is not the same will
	// Example

	src1 := []int{1, 2, 3}
	dest1 := make([]int, 1) // NOTE: des1 is a slice with one element. if the destination slice has length is zero, nothing will be copied
	copied := copy(dest1, src1)
	fmt.Printf("src1 is %v dest1 is %v copied is %v\n", src1, dest1, copied)
	// NOTE: Because dest1 has one element, it will only contain the first element of src1 and copied will return 1
}
