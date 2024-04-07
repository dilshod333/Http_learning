// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Enter a sentence:")
// 	reader := bufio.NewReader(os.Stdin)
// 	sentence, _ := reader.ReadString('\n')

// 	sentence = strings.TrimSpace(sentence)
// 	fmt.Print(sentence)
// 	words := strings.Fields(sentence)

// 	wordFreq := make(map[string]int)
// 	fmt.Println(wordFreq)

// 	for _, word := range words {

// 		word = strings.ToLower(word)
// 		wordFreq[word]++
// 	}

// 	fmt.Println("Word Frequencies:")
// 	for word, freq := range wordFreq {
// 		fmt.Printf("%s: %d\n", word, freq)
// 	}
// }
// EXERCISE 2

package main

import "fmt"

func singleNumber(nums []int) int {

	numFreq := make(map[int]int)
	for _, num := range nums {
		numFreq[num]++
	}
	for index, num := range numFreq {

		if num == 1 {
			return index
		}

	}
	return -1

}

func main() {

	num2 := []int{2, 3, 45, 6, 7, 3, 2}
	a := singleNumber(num2)
	fmt.Println(a)

}
