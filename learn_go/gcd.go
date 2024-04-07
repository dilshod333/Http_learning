package main

import "fmt"

func main() {
	n := 0
	// fmt.Scanln(&n)
	res := 1
	n2 := 0
	i := 1
	fmt.Scanln(&n)
	fmt.Scanln(&n2)
	if n > n2 {
		for i <= n2 {
			if n%i == 0 && n2%i == 0 {
				res = i
			}
			i += 1
		}
	} else {
		for i <= n {
			if n%i == 0 && n2%i == 0 {
				res = i
			}
			i += 1
		}
	}
	fmt.Println(res)
}
