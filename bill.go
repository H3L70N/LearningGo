package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string, items map[string]float64, tip float64) bill {
	bill := bill{
		name:  name,
		items: items,
		tip:   tip,
	}
	return bill
}

func (b bill) format() string {
	fs := fmt.Sprintf("Bill for %v \nIncludes the items: \n", b.name)
	total := 0.0
	for k, v := range b.items {
		fs += fmt.Sprintf(" %-25v %v$ \n", k, v)
		total += v
	}

	fs += fmt.Sprintf("%-25v %0.2f$", "Total:",total)
	return fs
}
