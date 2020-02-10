package main

import "fmt"

func isPifagor(a, b, c int) string {
	if a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b {
		return "Yes"
	}
	return "No"
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	result := isPifagor(a, b, c)
	fmt.Println(result)

}
