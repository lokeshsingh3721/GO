package main

import "fmt"

type address struct {
	state   string
	city    string
	country string
}

type details struct {
	name  string
	age   int
	place address
}

func readData(person details) {
	fmt.Printf("name is %s and age is %d and address is %s", person.name, person.age,person.place.state)
}

func main() {
	person := details{
		name: "lokesh",
		age:  20,
		place: address{
			state:   "haryana",
			city:    "faridabad",
			country: "india",
		},
	}

	readData(person)
}
