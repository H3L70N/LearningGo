package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	bill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
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

	fs += fmt.Sprintf("%-25v %0.2f$", "Subtotal:", total)
	fs += fmt.Sprintf("\n%-25v %0.2f$", "Tip:", b.tip)
	fs += fmt.Sprintf("\n%-25v %0.2f$", "Total:", total+b.tip)
	return fs
}

func (b *bill) insertTip(m float64) {
	(*b).tip = m
}

func (b *bill) insertItems(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
