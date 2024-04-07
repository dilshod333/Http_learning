// // package main

// // import "fmt"

// // func singleNumber(nums []int) int {
// // 	var array []int
// // 	for i := 0; i < len(nums); i++ {
// // 		unique := true
// // 		for j := i + 1; j < len(nums); j++ {
// // 			if nums[i] == nums[j] {
// // 				unique = false
// // 				break
// // 			}
// // 		}
// // 		if unique {

// // 			array = append(array, nums[i])
// // 		}
// // 	}

// // 	return array[0]
// // }

// // func main() {
// // 	var n int
// // 	fmt.Println("Enter the number of elements:")
// // 	fmt.Scanln(&n)

// // 	var nums []int
// // 	fmt.Println("Enter the elements:")
// // 	for i := 0; i < n; i++ {
// // 		var num int
// // 		fmt.Scan(&num)
// // 		nums = append(nums, num)
// // 	}

// // 	fmt.Println("Single Number:", singleNumber(nums))
// // }

// package main

// import "fmt"

// func singleNumber(nums []int) []int {
// 	var array []int
// 	for i := 0; i < len(nums); i++ {
// 		unique := true
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				unique = false
// 				break
// 			}
// 		}
// 		if unique {
// 			array = append(array, nums[i])
// 		}
// 	}

// 	return array
// }

// func main() {
// 	var n int
// 	fmt.Println("Enter the number of elements:")
// 	fmt.Scanln(&n)

// 	var nums []int
// 	fmt.Println("Enter the elements:")
// 	for i := 0; i < n; i++ {
// 		var num int
// 		fmt.Scan(&num)
// 		nums = append(nums, num)
// 	}

//		result := singleNumber(nums)
//		if len(result) == 1 {
//			fmt.Println("Single Number:", result[0])
//		} else {
//			fmt.Println("Multiple Single Numbers:", result)
//		}
//	}
package main

import "fmt"

func singleNumber(nums []int) int {
	var array []int
	for i := 0; i < len(nums); i++ {
		unique := true
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] == nums[j] {
				unique = false
				break
			}
		}
		if unique {
			array = append(array, nums[i])
		}
	}

	if len(array) > 0 {
		return array[0]
	}
	return -1
}
func main() {
	var n int
	fmt.Println("Enter the number of elements:")
	fmt.Scanln(&n)

	var nums []int
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	single := singleNumber(nums)
	if single != -1 {
		fmt.Println("Single Number:", single)
	} else {
		fmt.Println("No single number found")
	}
}
