// package main

// import "fmt"

// func main() {

// 	var myMap = map[string]int{"Dilshod": 1, "salim": 11, "vali": 15}
// 	fmt.Println(myMap)

// 	var m = map[int]
// 	name := "Dilshod"

// 	for _, n := range name {
// 		char := string(n)
// 		fmt.Println(char)
// 	}

// 	type people struct {
// 		name string
// 		age  int
// 	}

// 	surname := "DilshaAdbek"

// 	for i := 0; i < len(surname); i++ {
// 		fmt.Println(surname[i])
// 	}

// }

// // palindrom

//	func findPalindrome(n){
//		for i := 0; i < len(n); i++{
//			for j := i + 1; j < len(n); j++{
//				substr[i:j]
//				i
//			}
//		}
//	}
package main

import "fmt"

func firstUniqChar(s string) int {
	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}


	for i, char := range s {
		if count[char] == 1 {
			return i
		}
	}

	return -1 
}

func main() {
	s := "somsa"
	index := firstUniqChar(s)
	fmt.Println(index)
}
