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

func ( details) checkAge(age int) bool {
	if age>=18 {
		return true
	}
	return false
}

func main() {
	person := details{
		name: "lokesh",
		age:  10,
		place: address{
			state:   "haryana",
			city:    "faridabad",
			country: "india",
		},
	}

	var isAge = person.checkAge(person.age)

	fmt.Println(isAge)

	readData(person)
}
