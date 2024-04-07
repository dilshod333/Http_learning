// // // package main

// // // import (
// // // 	"fmt"
// // // 	"strings"
// // // 	// "strings"
// // // )

// // // func main() {
// // // 	s := make(map[string]int)
// // // 	n := "Dilshod dilmurodov hey qaley"
// // // 	a := strings.Fields(n)
// // // 	for _, key := range a {
// // // 		s[key]++

// // // 	}
// // // 	fmt.Println(s)

// // // }
// // package main

// // import (
// // 	"fmt"
// // 	"strings"
// // )

// //	func main() {
// //		s := make(map[string]int)
// //		n := "Dilshod dilmurodov hey qaley"
// //		a := strings.Fields(n)
// //		for _, key := range a {
// //			s[key]++
// //		}
// //		fmt.Println(s)
// //	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	s := make(map[string]int)

// 	fmt.Println("Enter a sentence: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')

// 	words := strings.Fields(input)

// 	for _, word := range words {
// 		s[word] = len(word)
// 	}

// 	fmt.Println("Count each word: ")
// 	for word, count := range s {
// 		fmt.Printf("%s: length %d\n", word, count)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	m := make(map[int]int)

// 	slice := []int{}

// 	n := [5]int{2, 3, 4, 5, 6}

// 	for _, num := range n {
// 		m[num] = num * 2

// 		slice = append(slice, m[num])
// 	}

// 	fmt.Println(m, slice)
// }

/*
func removeElement(slice []int, element int) []int {
    for i, v := range slice {
        if v == element {
            return append(slice[:i], slice[i+1:]...)
*/

package main

import "fmt"

func main() {
	var a int = 0
	fmt.Println(b)
	fmt.Println(a)
}
