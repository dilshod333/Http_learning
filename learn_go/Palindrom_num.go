package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	org := n

	reversed := 0
	a := 0
	for n != 0 {
		a = n
		a %= 10
		reversed = reversed*10 + a
		n = int(n / 10)
	}
	if reversed == org {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
