package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello my friends"
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	ages := []int{10, 20, 13, 14, 12, 18, 29, 30, 24, 47}
	sort.Ints(ages)
	fmt.Println(ages, sort.SearchInts(ages, 30))
}
