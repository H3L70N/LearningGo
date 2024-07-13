package main

import "fmt"

func updateValue(n *int) {
	*n = 10
}

func main() {
	number := 5
	pointer := &number
	updateValue(pointer)

	fmt.Println(number)
	// fmt.Println("Memory address of number", &number)
}
