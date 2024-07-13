package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	names := strings.Split(strings.ToUpper(n), " ")
	var initials []string
	for _, v := range names {
		initials  = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	} else if len(initials) == 0 {
		return "", ""
	} else {
		return initials[0], ""
	}
}

func main() {
	fmt.Println(getInitials("helton guambe"))
}
