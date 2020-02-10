package main

import "fmt"

// Solve ...
func Solve(a, b, c int) string {
	var result string
	if a == 0 {
		result = Linear(b, c)
	} else {
		result = Quadratic(a, b, c)
	}
	return result
}

//Linear ...
func Linear(b, c int) string {
	if b == 0 {
		return "Has No roots"
	}
	return "Has 1 root"
}

// Quadratic ...
func Quadratic(a, b, c int) string {
	D := b*b - 4*a*c
	if D == 0 {
		return "Has 1 root"
	} else if D > 0 {
		return "Has 2 roots"
	}
	return "Has No roots"

}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(Solve(a, b, c))
}
