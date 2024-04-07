package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Println("Enter the number: ")
	count := 0
	fmt.Scanln(&a)
	var array []int
	for i := 0; i < a; i++ {
		num := 0
		fmt.Scanln(&num)
		array = append(array, num)
		count += num

	}
	fmt.Println(array, count)
}
