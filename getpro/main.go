package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Equation ...
type Equation struct {
	A             int64  `json:"a"`
	B             int64  `json:"b"`
	C             int64  `json:"c"`
	TotalEquation string `json:"total_equation"`
	Result        string `json:"result"`
}

// getInfo ...
func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	a, _ := strconv.ParseInt(params["a"], 10, 64)
	b, _ := strconv.ParseInt(params["b"], 10, 64)
	c, _ := strconv.ParseInt(params["c"], 10, 64)

	answerEquation := fmt.Sprintf("%dx^2 + %dx + %d = 0", a, b, c)
	result := Solve(int(a), int(b), int(c))

	e := Equation{
		A:             a,
		B:             b,
		C:             c,
		TotalEquation: answerEquation,
		Result:        result,
	}
	json.NewEncoder(w).Encode(e)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{a}/{b}/{c}", getInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

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
