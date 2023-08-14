package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Arithmatic in Go")
	// adding
	x, y := 5, 4
	fmt.Println(x + y)

	// multiplying
	fmt.Println(x * y)

	// division

	fmt.Println(x / y)

	// subtraction
	fmt.Println(x - y)

	// modulus
	fmt.Println(x % y)

	// exponent
	fmt.Println(x ^ y)

	fmt.Println(x & y)

	fmt.Println(math.Pow(float64(x), float64(y)))

	var a uint8 = 255
	a = a + 1
	fmt.Println(a)
	fmt.Println(string(rune(99)))
}
