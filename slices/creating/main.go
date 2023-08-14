package main

import "fmt"

func main() {
	// create an uninitialized slice
	var cities []string
	fmt.Println("cities is equal to ", cities == nil)
	fmt.Printf("cities is %#v\n", cities)

	// NOTE: Creating a slice using the make generic function

	nums := make([]int, 10)
	fmt.Printf("nums is %#v\n", nums)

	// NOTE: the s below is called a slice literal
	s := []string{"a", "b", "c"}
	s = append(s, "d")
	l := len(s)
	fmt.Println(l)
	fmt.Printf("s is of type: %T\n", s)

	// NOTE: the type names below is just a slice type of strings
	type names []string
	friends := names{"Dan", "Mike"}
	fmt.Println(friends)
	fmt.Printf("friends is of type %T\n", friends)

	// NOTE: the above can be done using:

	names_slice := []string{"Dan", "Mike"}
	friends2 := names_slice

	fmt.Printf("friends is of type %T\n", friends2)
}
