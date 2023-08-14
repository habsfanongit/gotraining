package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

func term_size() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		return 0, 0
	}
	return width, height
}

func main() {
	width, _ := term_size()

	fmt.Println(strings.Repeat("#", width))
	// play_ground()
	country := "țară"
	fmt.Println(len(country))

	fmt.Println("Byte (NOT RUNE) as position 1 of country is ", country[0])

	// lopp through the string

	for i := 0; i < len(country); i++ {
		fmt.Printf("%c\n", country[i])
	}

	for _, v := range country {
		fmt.Printf("%c", v)
	}
	fmt.Println("")

	// NOTE: Manually converting all string runes to UTF-8 is done by usng the utf8 package

	fmt.Println("Manual Implementation")
	for i := 0; i < len(country); {

		r, size := utf8.DecodeRuneInString(country[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println("")

	fmt.Println(strings.Repeat("#", width))
}

func play_ground() {
	var s string = "Hello"
	sl := len(s)
	b := []byte(s)
	sc := cap(b)

	fmt.Printf("Length is %d, b is %v and sc is %d\n", sl, b, sc)
	r := '🐍'
	snake := string(rune(r))

	fmt.Printf("r is %v\n", r)
	fmt.Println("Snake is ", snake)
	fmt.Printf("Character: %c\n", r)
	ss := fmt.Sprintf("%c\n", r)
	fmt.Println("ss is ", ss)
}
