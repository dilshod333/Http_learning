// // // // package main

// // // // import "fmt"

// // // // // func change(son *int, son2 *int) *int {
// // // // // 	a := *son
// // // // // 	*son = *son2
// // // // // 	*son2 = a

// // // // // 	fmt.Println(son, son2)
// // // // // }

// // // // // func main() {
// // // // // 	number := 10
// // // // // 	number2 := 12

// // // // // 	fmt.Println("Before changing numbers: ", number, number2)

// // // // // 	change(&number, &number2)

// // // // // 	fmt.Println("After changing numbers: ", number, number2)

// // // // // }

// // // // /*
// // // // Second exercises

// // // // */

// // // // func main() {
// // // // 	array := [...]int{1, 2, 3, 4, 5}
// // // // 	doubleIt(&array)
// // // // }

// // // // func doubleIt(*array []int) {
// // // // 	for i := 0; i < len(*array); i++ {
// // // // 		*array[i] *= 2
// // // // 	}
// // // // 	fmt.Println(*array)
// // // // }

// // // // package main

// // // // import "fmt"

// // // // func main() {
// // // // 	array := [...]int{1, 2, 3, 4, 6}
// // // // 	DoubleIt(&array)
// // // // 	fmt.Println(array)
// // // // 	fmt.Printf("Type %T", array)
// // // // }

// // // // func DoubleIt(array *[5]int) {
// // // // 	for i := 0; i < len(*array); i++ {
// // // // 		((*array)[i]) *= 2
// // // // 	}
// // // // 	fmt.Println(*array)
// // // // 	fmt.Println(array)
// // // // }

// // // package main

// // // import "fmt"

// // // func main() {

// // // 	number := 18
// // // 	number2 := 22
// // // 	swap(number, number2)
// // // 	fmt.Println("After changing: ", number, number2)

// // // 	swap2(&number, &number2)
// // // 	fmt.Println(number, number2)

// // // 	array := [...]int{2, 43, 4, 5, 6}
// // // 	doubleIt(&array)
// // // 	fmt.Println("After changing: ", array)

// // // 	slice := []int{234, 3, 4, 5, 5}
// // // 	sum(&slice)
// // // }

// // // func swap(a, b int) {
// // // 	temp := a
// // // 	a = b
// // // 	b = temp
// // // 	fmt.Println(a, b)
// // // }

// // // func swap2(a, b *int) {
// // // 	temp := *a
// // // 	*a = *b
// // // 	*b = temp
// // // 	fmt.Println(*a, *b)

// // // }

// // // func doubleIt(array *[5]int) {
// // // 	var slice []int
// // // 	for _, num := range *array {
// // // 		num *= 2
// // // 		slice = append(slice, num)
// // // 	}
// // // 	fmt.Println(slice)
// // // }

// // // func sum(slice *[]int) {
// // // 	slice2 := []int{}
// // // 	fmt.Println(slice2)
// // // 	s := 0
// // // 	m := 1
// // // 	for _, num := range *slice {
// // // 		s += num
// // // 		m *= num
// // // 		slice2 = append(slice2, s, m)
// // // 	}
// // // 	fmt.Println("Sum is: ", s)
// // // 	fmt.Println("MUltiply is: ", m)
// // // 	fmt.Println("Slice2: ", slice)
// // // }

// // package main

// // import "fmt"

// // func main() {
// // 	slice := []int{}
// // 	up(&slice)
// // }

// // func up(a *[]int){
// // 	m := *a
// // 	if len(m) == 0 {
// // 		fmt.Println("Slice is empty")
// // 		return
// // 	}
// // 	max := m[0]
// // 	for _, num := range m {
// // 		if num > max {
// // 			max = num
// // 		}
// // 	}
// // 	fmt.Println("Max number is:", max)
// // }

// // package main

// // import "fmt"

// // func main() {
// // 	array := []int{2, 33, 4, 5, 8}
// // 	fmt.Println("Original array is: ", array)
// // 	fmt.Println("\n")
// // 	update(&array)
// // 	fmt.Println("After changing: ", array)
// // }

// // func update(array *[]int) {
// // 	for i := 0; i < len(*array); i++ {
// // 		(*array)[i] += 10
// // 	}
// // 	fmt.Println(*array)
// // }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	n := ""

// 	fmt.Scanln(&n)

// 	check := &n

// 	res := ""
// 	for i := len(*check) - 1; i > -1; i-- {
// 		res += string((*check)[i])
// 	}
// 	fmt.Println(res)

// 	a := ""
// 	fmt.Scanln(&a)
// 	b := ""
// 	fmt.Scanln(&b)
// 	swap_char(&a, &b)
// 	fmt.Println(a, b)
// }

// func swap_char(a *string, b *string) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// 	fmt.Println("After changing: ", *a, *b)
// }

package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Println("Before swap:", a, b)

	swap(&a, &b)

	fmt.Println("After swap:", a, b)
}

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}


