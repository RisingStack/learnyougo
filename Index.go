package main

import (
	"fmt"
	"net/http"
)

// Index handles GET "/"
func Index(w http.ResponseWriter, r *http.Request) {
	index, err := Asset("data/index.tmpl")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(index))
}
