package main

import (
	"fmt"
)

func main() {
	sentence := "hello"

	for index, char := range sentence {
		a := string(char)
		fmt.Println(index, a)
	}

	fmt.Println(sentence)
}
