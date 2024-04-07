package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	a := 0
	b := 1
	c := a + b
	i := 2
	for i < n {
		a = b
		b = c
		c = a + b
		i += 1
	}
	fmt.Println(c)
}
