package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args ", os.Args)
	fmt.Println("Path", os.Args[0])
	fmt.Println("First Arg", os.Args[1])
	fmt.Println("Second Arg", os.Args[2])
	fmt.Println("Number of Args", len(os.Args))

	if first, err := strconv.ParseInt(
		os.Args[1], 0, 64); err == nil {
		fmt.Printf("The first Arg is %d and it's type is %T\n", first, first)
		// TODO: Fix this shit
	}
}
