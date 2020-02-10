package main

import (
	"fmt"
	"log"
	"net/http"
)

// getPing ...
func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello form Ping!")
}

// getPong ...
func getPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello form Pong!")
}

func main() {
	http.HandleFunc("/ping", getPing)
	http.HandleFunc("/pong", getPong)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
