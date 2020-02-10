package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
