// package api

// import (
// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// type Item struct {
// 	id   uuid.UUID `json: "id"`
// 	Name string    `json: "name"`
// }

// type Server struct {
// 	*mux.Router

// 	shoppingItems []Item
// }

// func NewServer() *Server {
// 	s := &Server{
// 		Router:        mux.NewRouter(),
// 		shoppingItems: []Item{},
// 	}
// 	return s
// }

package main

import "fmt"

var age int = 18

func main() {
	fmt.Println(age)
}
