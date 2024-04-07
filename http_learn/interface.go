package main

import (
	"fmt"
	"net/http"
	"strings"
)
func name(){

}
func main() {
	a := 8
	b := &a
	a = 11
	c := 8
	
	fmt.Println(a, b, &a, *b, &b)

}

