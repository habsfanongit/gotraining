package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello Go")
	fmt.Println("One more time")

	// variable swapping
	var a, b int
	a, b = 1, 2
	b, a = a, b

	fmt.Println(
		a, b,
	)
	var x, y float64
	x, y = 1.0, 2.0

	fmt.Printf("The value of x is %.0f	and y as %.0f\n", x, y)
	// j, i := 1, 0
	// fmt.Println(j / i)
	const p = 5
	const q = 6.5
	fmt.Printf("The type of p is %T and q is %T\n", p, q)
}
