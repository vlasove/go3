package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == 0 {
		if b == 0 {
			fmt.Println("Has No roots")
		} else {
			fmt.Println("Has 1 root")
		}
	} else {
		D := b*b - 4*a*c
		if D == 0 {
			fmt.Println("Has 1 root")
		} else if D > 0 {
			fmt.Println("Has 2 roots")
		} else {
			fmt.Println("Has No roots")
		}
	}
}
