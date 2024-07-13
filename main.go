package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Printf("Hello %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cicleFunction(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * math.Pow(r, 2)
}

func main() {
	nameOne := "Helton"
	nameTwo := "Alceu"
	nameThree := "Hidilson"

	names := []string{nameOne, nameTwo, nameThree}

	sayGreetings(nameOne)
	sayBye(nameTwo)
	cicleFunction(names, sayGreetings)
	cicleFunction(names, sayBye)
	fmt.Printf("%0.2f \n",circleArea(math.Sqrt(10)))
}
