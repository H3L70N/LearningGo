package main

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill (name string, items map[string]float64, tip float64) bill {
	bill := bill{
		name: name,
		items: items,
		tip: tip,
	}
	return bill
}