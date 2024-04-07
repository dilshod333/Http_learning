package main

import (
	"net/http" // all stuf about http that is why we imorted it
)

func main() {
	http.HandleFunc("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world")) // w --> response
	})

	http.ListenAndServe(":8080", nil) // http server

}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello word")

// }

// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe("", nil)
// }
