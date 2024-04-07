package main

import "fmt"

func main() {
	var slice []int
	for i := 0; i < 5; i++ {
		num := 0
		fmt.Scanln(&num)
		slice = append(slice, num)
	}
	var result []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != j && slice[i] == slice[j] {
				result = append(result, slice[i])
			}
		}
	}
	fmt.Println(result)

}
