package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func greet(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go greet("Hi")
	go greet("Hello")

	fmt.Println("Finish!")
	wg.Wait()
}
