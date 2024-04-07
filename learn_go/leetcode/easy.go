package main

import "fmt"

func singleNumber(nums []int) int {
	var array []int
	for i := 0; i < len(nums); i++ {
		unique := true
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				unique = false
				break
			}
		}
		if unique {
			array = append(array, nums[i])
		}
	}

	return array[0]
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

	fmt.Println("Single Number:", singleNumber(nums))
}
