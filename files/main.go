package main

import (
	"bufio"
	f "fmt"
	"os"
)

func main() {
	f.Println("J")
	buffer, _ := os.ReadFile("./main.go")
	content := string(buffer)
	f.Println(content)

	f.Println("Reading lines")

	readfile, _ := os.Open("./main.go")
	defer readfile.Close()

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		f.Println(fileScanner.Text())
	}
}
