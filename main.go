package main

import "fmt"

func genericFunc[T any](value T) T {
	return value
}

func main() {
	stringValue := genericFunc("Hello") // return type is string
	intValue := genericFunc(24)         // return type is int

	fmt.Println(stringValue)
	fmt.Println(intValue)
}
