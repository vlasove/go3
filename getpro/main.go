package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Person ...
type Person struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int64  `json:"age"`
}

// getInfo ...
func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	p := Person{
		Name:     params["name"],
		Lastname: params["lastname"],
	}
	age, err := strconv.ParseInt(params["age"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	p.Age = age
	json.NewEncoder(w).Encode(p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}/{lastname}/{age}", getInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
