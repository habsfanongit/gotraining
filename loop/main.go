package main

import "fmt"

func main() {
	// NOTE: Basic for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// NOTE: emulate while loop

	// j := 10
	// for j >= 0 {
	// 	fmt.Println(j)
	// 	j--
	// }

	// NOTE: an infinit loop

	// sum := 0
	// for {
	// 	sum++
	// 	fmt.Println(sum)
	// }
	// NOTE: using continue

	// for i := 0; i < 10; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// NOTE: true in the below loop makes it an infinit loop but it exits
	// using the break statement

	// count := 0
	// for i := 0; true; i++ {
	// 	if i%13 == 0 {
	// 		fmt.Printf("%d is divisible by 13 \n", i)
	// 		count++
	// 	}
	// 	if count == 10 {
	// 		break
	// 	}
	// }
	// fmt.Println("Loop completed")
	// NOTE: using labels

	names := [5]string{"Joshua", "Mark", "Ellen", "Larry", "Sonja"}
	friends := [2]string{"Carlos", "Joshua"}

	// NOTE: outer is a label

outer:
	for index, name := range names {
		for _, friend := range friends {
			fmt.Println(name)
			if name == friend {
				fmt.Printf("Found a friend %q at %d of the names array\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("After Label executed")
}
