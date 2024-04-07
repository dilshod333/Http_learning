// package main

// import (
// 	"fmt"
// )

// func topKFrequent(nums []int, k int) []int {
// 	mostFreq := make(map[int]int)

// 	for _, num := range nums {

// 		mostFreq[num]++
// 	}

// 	slice := []int{}
// 	for key, value := range mostFreq {
// 		if key > 1 {
// 			a := max(value)
// 			slice = append(slice, a)
// 			delete(mostFreq, a)
// 		}
// 	}

// 	return slice
// }

// func main() {
// 	n := []int{1, 1, 1, 2, 2, 3}
// 	k := 2
// 	a := topKFrequent(n, k)
// 	fmt.Println(a)

// }

package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {

	mostFreq := make(map[int]int)

	for _, num := range nums {
		mostFreq[num]++
	}
	if len(mostFreq) == 1 {
		fmt.Println(0)
	}

	slice := make([]int, 0, len(mostFreq))

	for key, value := range mostFreq {
		if len(mostFreq) >= 2 {
			a := max(value)
			fmt.Println(a)
			slice = append(slice, key)
			delete(mostFreq, key)
		}
	}

	return slice[:k]
}

func main() {
	n := []int{5, 5}
	k := 1
	a := topKFrequent(n, k)
	fmt.Println(a)

	n1, n2 := 7, 8
	fmt.Println(n1, n2)
	n1, n2 = n2, n1
	fmt.Println(n1, n2)

}
