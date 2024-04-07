package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the words:")
	n, err := reader.ReadString('\n') 
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Print the input string
	fmt.Println(n)
}
