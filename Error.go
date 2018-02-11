package main

import (
	"fmt"
	"net/http"
)

// Error handles GET "/error"
func Error(w http.ResponseWriter, r *http.Request) {
	errorPage, err := Asset("data/error.tmpl")
	if err != nil {
		panic(err) // well, this is the error handler so I guess there's nothing better to do.
	}
	fmt.Fprint(w, string(errorPage))
}
