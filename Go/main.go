package main

import (
	"fmt"
)

func main() {
	var slice []int
	for i := 0; i < 5; i++ {
		num := 0
		fmt.Scanln(&num)
		slice = append(slice, num)

	}
	var slice2 []int
	sum := 0
	for i, num := range slice {
		if num > 0 {
			sum += num
			slice2 = append(slice2, i)
		}
	}
	if sum != 0 {
		fmt.Println(sum)
	} else {
		fmt.Println(-1)
	}

	// var slice
}

func subSequency(x []int, target int) []int {
	var result []int
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i]+x[j] == target {
				result = append(result, x[i], x[j])
			}
		}
	}
	return result
}
