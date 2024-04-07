package main

import "fmt"

func main() {
	n := 0
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	t := 0
	for i := 2; i < n; i++ {
		t = 0
		if n%i == 0 {
			break
		} else {
			t = 1
		}
	}
	if t == 1 {
		fmt.Println("This is the prime number: ", n)
	} else {
		fmt.Println("This is not the prime number: ", n)
	}
}
