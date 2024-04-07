package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	res := 0
	s := n
	fmt.Println(s)
	for i := 1; i < n; i++ {
		if s%i == 0 {
			fmt.Println(i)
			res += i

		}
	}

	if res == s {
		fmt.Println("perfect number ")
	} else {
		fmt.Println("not Perfect ")
	}

}
