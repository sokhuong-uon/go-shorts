package main

import "fmt"

func sum[T int | float64](slice []T) T {
	var sum T
	sum = 0

	for _, v := range slice {
		sum += v
	}

	return sum
}

func main() {
	intSlice := []int{3, 6, 1, 8}
	floatSlice := []float64{3.3, 6.1, 8.5}

	intSum := sum(intSlice)
	floatSum := sum(floatSlice)

	fmt.Println(intSum)
	fmt.Println(floatSum)
}
