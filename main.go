package main

import "fmt"

func main() {
	name := "Helton"
	items := map[string]float64{
		"Sandwich": 6.99,
		"Soda":     3.55,
	}
	tip := 0.0

	bill := newBill(name, items, tip)

	fmt.Println(bill)

}
