// // package main

// // import "fmt"

// // func change(x, y *int){
// // 	temp := *x
// // 	*x = *y
// // 	*y = temp
// // 	fmt.Println(*x, *y)
// // 	return *x, *y
// // }

// // func main() {
// // 	// pointer
// // 	//map

// // 	a := 8
// // 	x := &a
// // 	*x = 9
// // 	// fmt.Println(&a, x)

// // 	n := 1
// // 	n2 := 3
// // 	fmt.Printf("Before changing: n= %d and n2 = %d\n", n, n2)
// // 	change(&n, &n2)
// // 	fmt.Printf("After changing: n= %d and n2= %d \n", n, n2)
// // }

// /*
// Pointer to do

// */

// package main

// import "fmt"

// func main() {
// 	n := 0
// 	fmt.Println("Enter how many number you want to enter: ")
// 	fmt.Scanln(&n)
// 	slice := []int{}
// 	for i := 0; i < n; i++ {
// 		num := 0
// 		fmt.Println("Enter the number: ", i+1)
// 		fmt.Scanln(&num)
// 		slice = append(slice, num)
// 	}
// 	numMap := make(map[int]int)

// 	for _, num := range slice {
// 		numMap[num]++
// 	}
// 	fmt.Println(numMap)
// }

// character counter

package main

import (
	"fmt"
)

func main() {
	input := ""
	fmt.Scanln(&input)

	charFreq := make(map[string]int)

	for _, char := range input {

		charStr := string(char)
		fmt.Println(charStr)
		charFreq[charStr]++
	}

	fmt.Println(charFreq)

	a := vowelCount(input)
	fmt.Println(a)

	var My_value_2 string = "Dilshodbek nima gap"
	fmt.Println(My_value_2)
}

// vowel counter

func vowelCount(a string) int {
	count := 0
	n := "aouie"

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(n); j++ {
			if a[i] == n[j] {
				count++
				break
			}
		}

	}
	fmt.Println(count)
	return count
}
