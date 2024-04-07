package main

import "fmt"

type Any struct {
	Make, Model, Color string
	Year, Weight       int
}

func main() {
	// var person Person
	var s string = "Привет, мир"
	fmt.Println(len(s))

	type Person struct {
		name        string
		surname     string
		age         int
		phoneNumber string
	}

	person := Person{
		name:        "Dilshod",
		surname:     "Dilmurodov",
		age:         18,
		phoneNumber: "998909640820",
	}

	//struct ichida struct yaratish

	person2 := Person{
		name:        "Ali",
		surname:     "Valiyev",
		age:         19,
		phoneNumber: "998907654332",
	}

	person3 := Person{
		"Ali",
		"Valiyev",
		19,
		"998907654332",
	}

	// person3 := Person{
	// 	"vali",
	// 	"salim",
	// 	22,
	// 	"998907896756",
	// }

	fmt.Println(person.name, person.surname, person.age, person.phoneNumber)

	fmt.Println(person2.name, person2.surname, person2.age, person2.phoneNumber)

	fmt.Println(person3.age, person3.name, person3.phoneNumber)

	//   interfeyslar

	//   kodni kengaytirish

	type Travel interface {
		flyTicket() int
		tryTravel()
	}
	// var car Any

	car := Any{Make: "Nissan", Model: "GT", Color: "Black", Year: 2024, Weight: 456}
	fmt.Println(car)

}
