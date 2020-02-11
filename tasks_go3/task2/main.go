package main

import (
	"fmt"
	"strings"
)

func Pyramid(n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += strings.Repeat(" ", n-(i+1))
		result += strings.Repeat("@", (i*2)+1)
		result += "\n"
	}
	return strings.TrimRight(result, "\n")
}

func main() {
	fmt.Print(Pyramid(2))
}
