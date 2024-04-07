package main

import "fmt"

func main() {
	a := ""
	n := ""
	fmt.Println("Enter the word: ")
	fmt.Scanln(&n)

	for i := len(n) - 1; i > -1; i-- {
		a += string(n[i])
	}
	fmt.Println(a)

}
