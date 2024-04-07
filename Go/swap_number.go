// // package main

// // import "fmt"

// // func main() {
// // 	a := 5
// // 	b := 10

// // 	fmt.Println("Before swap:", a, b)

// // 	// swap(&a, &b)

// // 	fmt.Println("After swap:", a, b)

// // 	n := 12
// // 	var n2 *int = &n
// // 	fmt.Println(*n2)

// // 	*n2 = 10
// // 	fmt.Println(n, *n2)

// // 	var n3 *int
// // 	fmt.Println(n3)
// // }

// // // func swap(a, b *int) {
// // // 	*a = *a + *b
// // // 	*b = *a - *b
// // // 	*a = *a - *b
// // // }

// // Go program to illustrate the
// // concept of the Pointer to Pointer
// package main

// import "fmt"

// // Main Function
// func main() {

// 	var V int = 100

// 	var pt1 *int = &V

// 	var pt2 **int = &pt1
// 	**pt2 = 20

// 	fmt.Println(*pt1, pt2, *pt2, **pt2)

//		const s = 9
//		s = 11
//		fmt.Println(s)
//	}
//
// Go program to illustrate the
// concept of the call by value
