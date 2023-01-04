package main

import "fmt"

type UseRefReturnType[U any] struct {
	current U
}

func useRef[T any](value T) UseRefReturnType[T] {

	ref := UseRefReturnType[T]{
		current: value,
	}

	return ref
}

func main() {
	boolRef := useRef(false)
	stringRef := useRef("Hello")

	fmt.Println(boolRef.current)
	fmt.Println(stringRef.current)
}
