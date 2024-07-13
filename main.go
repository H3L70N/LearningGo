package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Soup":    4.99,
		"Pie":     7.99,
		"Salad":   3.99,
		"pudding": 3.55,
	}

	fmt.Println(menu, menu["pudding"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
}
