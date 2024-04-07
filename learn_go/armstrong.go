package main

import "fmt"

func main() {
	var array []int
	n := 0
	fmt.Scanln(&n)
	org := n
	count := 0
	for n != 0 {
		a := n
		a %= 10
		count += 1
		array = append(array, a)
		n = int(n / 10)

	}
	i := 0
	res := 0
	for i < count {
		digit := array[i]
		powerDigiit := 1
		j := 0
		for j < count {
			powerDigiit *= digit
			j += 1
		}
		res += powerDigiit

	}
	if res == org {
		fmt.Println("true")

	} else {
		fmt.Println("false")
	}
}
