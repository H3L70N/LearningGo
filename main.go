package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt + ": ")
	name, error := r.ReadString('\n')
	name = strings.TrimSpace(name)
	return name, error
}


func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=====Create a new Bill=====")
	name, _ := getInput(reader, "Name")
	bill := newBill(name)
	return bill
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput(reader, "Choose option \nA - Add Item\nS - Save bill\nT - Tip bill\n>")
	opt = strings.ToUpper(opt)
	for opt != "S" {
		switch opt {
		case "A":
			name, price := addItem(reader)
			b.insertItems(name, price)
		case "T":
			tip := addTip(reader)
			fmt.Println("your tip is",tip)
			b.insertTip(tip)
			fmt.Println("final tip is",b.tip)
		default:
			fmt.Println("Not a valid option")
		}
		opt, _ = getInput(reader, "Choose option \nA - Add Item\nS - Save bill\nT - Tip bill\n>")
		opt = strings.ToUpper(opt)
	}
	b.saveBill()
	fmt.Println("Saved")
	
}

func addItem(r *bufio.Reader) (string, float64) {
	fmt.Println("====Adding Item to the bill====")
	name, _ := getInput(r, "name")
	rp := true
	var value string
	var price float64
	for rp {
		var err error
		value, _ = getInput(r, "price")
		price, err = strconv.ParseFloat(value, 64)
		if err == nil {
			rp = false
		} else {
			fmt.Println("Invalid price")
		}
	}
	return name, price
}

func addTip (r *bufio.Reader) float64 {
	fmt.Println("====Adding Tip to the bill====")
	rp := true
	var value string
	var tip float64
	for rp {
		var err error
		value, _ = getInput(r, "tip")
		tip, err = strconv.ParseFloat(value, 64)
		if err == nil {
			rp = false
		} else {
			fmt.Println("Invalid tip")
		}
	}
	return tip
}

func main() {
	bill := createBill()
	promptOptions(&bill)
	// fmt.Println(bill.format())
}
