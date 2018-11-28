package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Numbers struct {
	Number int `json:"number"`
}

func printResult(key string, msg string) string {
	return "{\"" + key + "\":\"" + msg + "\"}"
}

func factorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var number Numbers

	res := 1

	err := json.NewDecoder(r.Body).Decode(&number)
	if err != nil {
		panic(err)

	}

	result := number.Number

	for j := 2; j <= result; j++ {
		res = res * j

	}

	s := strconv.Itoa(res)

	if err != nil {
		panic(err)
	}
	if number.Number < 0 {
		fmt.Fprintf(w, printResult("Result", "You can´t insert negative number!"))
	} else {
		fmt.Fprintf(w, printResult("Result", s))
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/factorial", factorial).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
