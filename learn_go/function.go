package main

import (
	"fmt"
)

func reverse(a int) int {
	var reversed int
	// fmt.Println(reversed)
	for a != 0 {
		n := a
		n %= 10
		fmt.Println(n)
		reversed = reversed*10 + n
		a = int(a / 10)
	}
	return reversed
}

func main() {
	var n int
	var n2 []int
	fmt.Scanln(&n)
	a := reverse(n)
	fmt.Println(a)

	fmt.Println(divide(n, a))
	for i := 0; i < 5; i++ {
		num := 0
		fmt.Println("Enter the number: ")
		fmt.Scanln(&num)
		n2 = append(n2, num)
	}
	fmt.Println(uniqNum(n2))
	palindromNumber()
	palindromNumber()
	palindromNumber()
	palindromNumber()

}

func divide(a int, b int) int {
	result := int(a / b)
	return result
}
func multiply(a int, b int, c int) int {
	return a * b * c

}
func minus(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func plus(a int, b int) int {
	return a + b
}

func palindromNumber() {
	fmt.Println("This is palindrome function")
}

func armstrong(n int) string {
	return "salom"
}

func uniqNum(n []int) int {
	// unique := true
	var slice []int
	for i := 0; i < len(n); i++ {
		unique := true
		for j := 0; j < len(n); j++ {
			if i != j && n[i] == n[j] {
				unique = false
				break
			}
		}
		if unique {
			slice = append(slice, n[i])
		}
	}
	fmt.Println(slice)
	if len(slice) >= 1 {
		return slice[0]
	} else {
		return -1
	}
}

// func count_char(n []string) int {
// 	fmt.Println("NAh")
// }
