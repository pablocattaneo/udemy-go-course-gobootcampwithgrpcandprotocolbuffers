package main

import "fmt"

func main() {
	for i := 1; i <= 1; i++ {
		println("Iteration:", i)
	}

	// iterating over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
