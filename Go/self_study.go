/*
flexible
concurrency
fast output
library
garbage collection


two different ways to declare variabe

"fmt--->formating page <---->  print insidi in it"


*/

package main

import "fmt"

func main() {
	var a, b int
	a, b = 19, 56
	// fmt.Println(a, b)
	a = 8
	b = 9
	fmt.Println(a, b)

	var d, e int
	d = 9
	e = 10
	fmt.Println(d, e)

	var c, f string
	c, f = "Dilshod", "dilmurodov"
	fmt.Println(c, f)
	n := true
	n = false
	fmt.Println(n)

	var f1, f2 int
	f1, f2 = 7, 3
	result := f1 / f2
	fmt.Println(result)
}
