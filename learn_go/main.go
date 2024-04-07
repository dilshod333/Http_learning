// package main

// import "fmt"

// func main() {
// 	// var username string = "Dilshod"
// 	// username = "salom"
// 	// fmt.Println(username)

// 	n := 0
// 	b := 0
// 	fmt.Scanln(&n)
// 	for n != 0 {
// 		a := n
// 		a %= 10
// 		b += a
// 		n = int(n / 10)
// 	}
// 	fmt.Println(b)

// }

// 2 Factorial

// package main

// import "fmt"

// func main() {
// 	a := 0
// 	fmt.Scanln(&a)
// 	i := 1
// 	fact := 1
// 	for i <= a {
// 		fact *= i
// 		i += 1
// 	}
// 	fmt.Println(fact)

// }

// 3 Reverse string

// package main

// import "fmt"

// func main() {
// 	a := ""
// 	fmt.Println("Enter the string: ")
// 	fmt.Scanln(&a)

// 	l := -1
// 	r := len(a) - 1
// 	n := ""
// 	for r > l {
// 		n += string(a[r])
// 		r -= 1
// 	}
// 	fmt.Println(n)
// }

// 4-armstrong number checker

package main

import "fmt"

func main() {
	a := 0
	count := 0
	var slice []int
	fmt.Println(slice)
	fmt.Scanln(&a)
	org := a
	for a != 0 {
		n := a
		n %= 10
		count += 1
		slice = append(slice, n)
		a = int(a / 10)
	}
	i := 0
	res := 0
	for i < len(slice) {
		a := slice[i]
		powerDigit := 1
		j := 0
		for j < count {
			powerDigit *= a
			j += 1
		}
		res += powerDigit
		i += 1
	}
	fmt.Println(res)
	if org == res {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	fmt.Println("%T", org)
}
