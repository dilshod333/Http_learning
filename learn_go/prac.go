// package main

// import "fmt"

// func main() {
// 	n := 0
// 	var a []int
// 	println("hey enter the size: ")
// 	fmt.Scanln(&n)
// 	for i := 0; i < n; i++ {
// 		num := 0
// 		println("Enter the number: ")
// 		fmt.Scanln(&num)
// 		a = append(a, num)
// 	}
// 	println(a)
// }

package main

import "fmt"

// func main() {
// 	a := 0
// 	s := 0
// 	d := 0
// 	for i := 0; i < 2; i++ {
// 		fmt.Println("Enter the number: ", d+1)
// 		fmt.Scanln(&a)
// 		s += a
// 		d = 1
// 	}
// 	fmt.Println("sum is: ", s)

// }
func main() {
	a := 0
	s := 1
	fmt.Println("Enter the number: ")
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {
		s *= i
	}
	fmt.Println(s)
}
