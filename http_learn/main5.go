package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	brand string
	year  int
}

var car []Car

func carHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCar(w, r)
	case http.MethodPost:
		postCar(w, r)

	default:
		http.Error(w, "invalid ", http.StatusMethodNotAllowed)
	}
}

func getCar(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(car)

	fmt.Fprint(w, "Get method useed %v", car)
}

func postCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	err := json.NewDecoder(r.Body).Decode(&car)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Post new person ", car)
}

func main() {

	http.HandleFunc("/car", carHandler)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello does anybody here"))
	})

	http.HandleFunc("/", helllo)

	http.ListenAndServe(":8080", nil)

}
func helllo(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Main page</h1>")

	case "/hello/":

		fmt.Fprint(w, "This is hello page ")
	case "/any/":
		fmt.Fprint(w, "<h1>This is any page</h1>")
	}

}
