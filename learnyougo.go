package main

import (
	"log"
	"net/http"
)

type example struct {
	Name string
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/exercises", Exercises)
	http.HandleFunc("/instructions", Instructions)
	http.HandleFunc("/error", Error)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
