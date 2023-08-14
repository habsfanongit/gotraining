package main

import "fmt"

func main() {
	// arrays are fixed size in Go
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// Slices are dynamic size in Go
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	// adding items to a slice
	i := 6
	slice = append(slice, i)
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)

	// Maps are dynamic size in Go and they are like dictionaries
	// m variable below has a key of type string and a value of type int
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	// struct are like classes in OOP languages
	// struct is delclared with the keywords struct
	var s struct {
		name string
		age  int
	}
	s.name = "Sammy"
	s.age = 20
	fmt.Println(s)

	// using struct as type
	type person struct {
		name string
		age  int
	}
	p := person{"Sammy", 20}
	fmt.Println(p)

	// pointer in Go
	x := 5
	ptr := &x
	fmt.Printf("the type of ptr is %T and the value is %v\n", ptr, ptr)

	// dereferencing the pointer
	fmt.Println(*ptr)
}
