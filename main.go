package main

import (
	"fmt"
	"time"
)

func greet(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go greet("Hi")
	greet("Hello")
}
