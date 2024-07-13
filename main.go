package main

import "fmt"

func main() {
	var ages [3]int = [3]int{20, 30, 40}

	var names = [4]string{"One", "Two", "Three", "Four"}

	scores := [4]float64{10.4, 234, 5.7, 6.8}
	scores[3] = 10.25

	var slice = []string{"Slice1", "Slice2", "Slice3"}
	slice = append(slice, "MoreSlice")


	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	fmt.Println(scores, len(scores))
	fmt.Println(slice, len(slice))

	rangeOne := names[1:3]

	fmt.Println(rangeOne)

}
