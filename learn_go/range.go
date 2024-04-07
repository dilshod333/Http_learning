// package main

// import "fmt"

// func main() {
// 	var slice []int
// 	var n int
// 	fmt.Println("Enter the length number: ")
// 	fmt.Scanln(&n)
// 	result := 0
// 	for i := 0; i < n; i++ {
// 		var num int
// 		fmt.Println("Enter the number: ", i)
// 		fmt.Scanln(&num)
// 		result += num
// 		slice = append(slice, num)

// 	}
// 	fmt.Println("The result is ", result)
// }

package main

import "fmt"

func main() {
	len := 0
	fmt.Scanln(&len)

	var slice []int
	for i := 0; i < len; i++ {
		num := 0
		fmt.Println("Enter the number: ", i)
		fmt.Scanln(&num)
		slice = append(slice, num)
	}
	n := 0
	for _, num := range slice {
		if num%2 == 0 {
			n += num
		}
	}
	fmt.Println(n)
	var slice2 []float32

	for i := 0; i < 5; i++ {
		f := 0
		fmt.Println("Enter the float number: ")
		fmt.Scanln(&f)
		slice2 = append(slice2, float32(f))

	}
	findIt := maxNum(slice2)
	fmt.Println(findIt)
}

func maxNum(a []float32) float32 {
	m := a[0]
	if len(a) >= 2 {
		for _, num := range a {
			if num > m {
				m = num
			}
		}
	}
	return m
}
