package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "index.html")
		if r.Method == "POST"{
			
		}
	case "/answer/":
		http.ServeFile(w, r, "answer.html")
	case "/info/":
		fmt.Fprint(w, "<h1>Here is nothing</h1>")

	}
}

func check(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		
	}
}

func main() {

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
