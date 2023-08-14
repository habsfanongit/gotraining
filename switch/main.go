package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: Traditional swich statement
	// Go does not need a break statememnt as the compiler ads it
	// even the default statement is NOT required but it's prefferd to use
	lang := "Go"
	switch lang {
	case "Go", "Golang":
		fmt.Println("You are learning Go")
	case "Python":
		fmt.Println("You are learning Python")
	default:
		fmt.Println("Default")
	}

	// NOTE: switch true compares a boolean expression
	// The true part of switch true can be ommitted
	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even Number ")
	case n%2 != 0:
		fmt.Println("Odd Number")
	}
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("Good morning")
	} else {
		fmt.Println("Good afternoon")
	}
}
