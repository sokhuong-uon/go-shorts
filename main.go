package main

import (
	"fmt"
)

func main() {
	fmt.Println("range:")
	arr := [6]int{2, 4, 6, 1, 8, 6}
	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Println("while")
	a := 0
	for a < 10 {
		fmt.Println(a)
		a += 2
	}

	fmt.Println("for")
	for b := 0; b < 10; b++ {
		fmt.Println(b)
	}
}
