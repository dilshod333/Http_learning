package main

import (
	"fmt"
	"time"
)

var a int = 8

func main() {
	fmt.Println(a)

	var a = "Dilshod"
	fmt.Println("after changing ", a)
	// async tarzda ishlab ketadi concurrency

	go cheeck("Salom !!!")     // go yozilgani uchun goroutine deyiladi   bir vartni uzida mustaqil bulib ishlatiladi
	go cheeck("Ikkinchi qism") // syncr tarzda ishlidi

	time.Sleep(2 * time.Second)

	fmt.Scanln()
}

func cheeck(input string) {
	for i := 0; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
