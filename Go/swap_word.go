package main

import "fmt"

func main() {
	word := ""
	fmt.Scanln(&word)
	w := &word

	result := ""
	for i := len(*w) - 1; i > -1; i++ {
		result += string((*w)[i])
	}

	fmt.Println(result)

}
