package main

import "fmt"

/*
*
* NOTE:
* When creating a slice in go, the elements of the slice are stored in a backing array, not in the slice itself.
* Go implelements a data structure called slice header to track the slice
* the silce header contains 3 fields
* 1. the address of the backing array (a pointer)
* 2. the length of the slice
* 3. the capacity of the slice.
* The slice header is the run time representation of a slice
* a **nil slice** does not have a backing array as it doesnt have any elements.
* NOTE: when a slice is created by an expression from an array or another slice, the backing array is shared between all slices
* See below example
* */
func main() {
	//	using_expressions()
	// usingAppen()
	// more_append()
	advanced_slices()
}

func advanced_slices() {
	//->NOTE: the example below will show how to retrieve elements from a slice
	//after its overwritten

	letters := []string{"A", "B", "C", "D", "E", "F"}
	// NOTE: Lets modify letters to {"A", "X", "Y"}

	letters = append(letters[:1], "X", "Y")

	// Print letters, its length and Capacity

	fmt.Printf("Letters is %v, It's length is %d and it's capacity is %d\n", letters, len(letters), cap(letters))
	// because letter's length is 3 the following line will fail with the following error.
	// panic: runtime error: index out of range [4] with length 3
	// BUG: Failing code letters[4] = "Z"

	// NOTE: even though the slice length is 3, we can retrieve "D", "E", "F" using the slice expression because it uses the slice's backing array.
	// using the following:

	fmt.Println(letters[3:6])
	// NOTE: the above line prints [D E F]
}

func more_append() {
	// NOTE: Append function internal examples

	// NOTE: create a nil slice

	var slice []int
	fmt.Printf("The value of slice is %#v\n", slice)
	fmt.Printf("The memory address of slice is %p\n", &slice)
	// NOTE: nil slices have 0 length and 0 capacity because they dont have a backing array

	fmt.Println("slice length is ", len(slice), " and capacity is ", cap(slice))
	slice = append(slice, 1, 2)
	fmt.Println("Value of slice is ", slice)
	// NOTE: Note the capacity and length after adding the two items

	fmt.Printf("slice length is %d, and it's capacity is %d \n", len(slice), cap(slice))
	// NOTE: the above code shows length and capacity of 2 because the compiler creates a new backing array when the current slice capcity is full.

	// lets add another item to the slice

	slice = append(slice, 3)

	fmt.Printf("slice length is %d, and it's capacity is %d \n", len(slice), cap(slice))
	// NOTE: the line above will show a length of three (1,2,3) and a capcity of four because the compiler
	// is attempting to avoid getting a full capacity if another append operation occurs
	// for example, the following operation will not modify the capacity because
	// it's adding one element to the slice which has a Capacity of 4 and three elements.

	slice = append(slice, 4)

	fmt.Printf("slice length is %d, and it's capacity is %d \n", len(slice), cap(slice))
	// NOTE: Now we have a slice with four elements and a capacity of four.
	// The next call to append, the compiler will create a capacity of 8 NOT 5.

	slice = append(slice, 5)

	fmt.Printf("slice length is %d, and it's capacity is %d \n", len(slice), cap(slice))
	// NOTE: From now on, once the Capacity is reached, the compiler will create a new slice using the current capacity multiplied by 2 (8*2)

	slice = append(slice, 6)
	// The above call will only add one item which does need increased capacity because we are at 8 already

	fmt.Printf("slice length is %d, and it's capacity is %d \n", len(slice), cap(slice))

	// NOTE: since we have 6 items and a capcity of 8, adding 3 more items will cause
	// the compiler to multiply the current capacity of 2. (8*2=16)

	slice = append(slice, 7, 8, 9)

	fmt.Printf("slice length is %d, and it's capacity is %d \n", len(slice), cap(slice))
}

func usingAppen() {
	// NOTE: Using append to modify a slice doesn't impact the backing array which means modification of the origibal slice
	// will not impact the new ones.
	german_cars := []string{"Benz", "BMW", "VolksWagen", "Audi", "Porsche"}
	sports_cars := []string{}
	fmt.Printf("Memory address of german_cars is %p\n", &german_cars)
	fmt.Printf("Memory address of sports_cars is %p\n", &sports_cars)
	sports_cars = append(sports_cars, german_cars[3:5]...)
	german_cars[4] = "Opel"
	fmt.Println("German Cars are ", german_cars, "Length is ", len(german_cars), "Capacity is ", cap(german_cars))
	fmt.Println("Sports Cars are ", sports_cars, "Length is ", len(sports_cars), "Capacity is ", cap(sports_cars))
	luxury_cars := []string{}
	luxury_cars = append(luxury_cars, german_cars[:3]...)
	fmt.Printf("Memory address of luxury_cars is %p\n", &luxury_cars)
	fmt.Println("luxury Cars", luxury_cars, "Length is ", len(luxury_cars), "Capacity is ", cap(luxury_cars))
}

func using_expressions() {
	arr_a := [4]int{10, 20, 30, 40}
	s_a, s_b := arr_a[0:2], arr_a[2:4]
	s_a[1] = 1000 // NOTE: after this change all slices and the original array are changed.
	fmt.Println("The original Array (arr_a)", arr_a)
	fmt.Println("s_a is ", s_a)
	fmt.Println("s_b is ", s_b)

	fmt.Printf("Memory address of arr_a is %p\n", &arr_a)
	fmt.Printf("Memory address of s_a is %p\n", &s_a)
	fmt.Printf("Memory address of s_b is %p\n", &s_b)

	fmt.Println("USING SLICES")
	// NOTE: The same behaviour can be observed if a slice is created from another slice using an expression
	s_c := []int{10, 20, 30, 40} // slice used as the source slice of s_d and s_e
	s_d, s_e := s_c[0:2], s_c[1:3]

	s_d[1] = 1000

	fmt.Println("s_c (Original) is ", s_c)
	fmt.Println("s_d is ", s_d)
	fmt.Println("s_e is ", s_e)

	fmt.Printf("Memory address of s_d is %p\n", &s_c)
	fmt.Printf("Memory address of s_d is %p\n", &s_d)
	fmt.Printf("Memory address of s_e is %p\n", &s_e)
}
