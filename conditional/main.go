package main

import (
	"fmt"
	"os"
)

func main() {
	debuggerShouldPauseHere := true
	fmt.Println(debuggerShouldPauseHere)

	arg1 := os.Args[1]
	fmt.Println(arg1)
}
