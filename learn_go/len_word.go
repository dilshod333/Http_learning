package main

import "fmt"

func main() {
	var slice []string

	for {
		var input string
		fmt.Print("Enter the string: ")
		_, err := fmt.Scanln(&input)

		if err != nil && input == "" {
			break
		}
		slice = append(slice, input)

	}
	fmt.Println(slice)
	if len(slice) >= 2 {
		max_len := len(slice[0])
		for _, n := range slice {
			if len(n) > max_len {
				max_len = len(n)
			}
		}
		fmt.Println(max_len)

	}
	var slice3 = [...]int{1, 4, 5, 6}
	removeDup(slice3[:])

}

func removeDup(elements []int) {
	for _, num := range elements {
		fmt.Print(num, " ")
	}

}
