package main

import (
	"fmt"
	"net/http"
)

checkHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path{
		fmt.Prit
	}
	if err != nil {
		fmt.Fprintf(w, "404 not found")
	}

	fmt.Fprintln(w, "Hey this is the response", resp)
}

func main() {
	http.HandlerFunc("/", checkHandler)

	http.ListenAndServe(":8080", nil)
}
