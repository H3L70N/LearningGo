package main

import "fmt"

func count(n int) {
	i := 1
	for i <= n {
		fmt.Println(i)
		i++
	}
}