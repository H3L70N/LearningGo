package main

import "fmt"

func main() {
	name := "Helton"
	items := map[string]float64{
		"Sandwich": 6.99,
		"Soda":     3.55,
		"Icecream": 2.45,
	}
	tip := 10.0



	bill := newBill(name)

	for k, v := range items {
		bill.insertItems(k, v)
	}

	bill.insertTip(tip)

	fmt.Println(bill)
	fmt.Println(bill.format())
}
