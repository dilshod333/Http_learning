package main

import "fmt"

func main() {
	// m := make(map[string]int{"Salom":14})
	m := make(map[string]int)

	fmt.Println(m)

	m2 := map[string]int{"ALi": 12, "Dilshod": 18}

	fmt.Println(m2)

	var m3 = make(map[string]int)

	fmt.Println(m3)

	words := make([]string, 8, 10)
	words[0] = "hi"
	words[1] = "the"
	fmt.Println(len(words), words)

	var odds = []int{1, 3}
	var evens = []int{8, 10}
	fmt.Println(append(odds, evens...))

	cheeses := [5]string{"parm", "blue", "gouda"}
	fmt.Println(cheeses[2], cheeses)

	numbers := map[string]int{
		"one":    1,
		"ten":    10,
		"twenty": 20,
	}
	delete(numbers, "twenty")
	fmt.Println(len(numbers), numbers)

	countryCodes := map[string]string{
		"93":  "Afghanistan",
		"374": "Armenia",
		"61":  "Australia",
	}
	country, err := countryCodes["61"]
	fmt.Println(country, err)

	if str, ok := weird(); ok {
		fmt.Println(str, "is a ghost")
	}

	d := (plusSeven(20))
	fmt.Println(d)
}
func weird() (string, bool) {
	return "casper", true
}

func plusSeven(num int) int {
	plusFive := func(num int) int {
		return num + 5
	}
	return plusFive(num) + 2
}
