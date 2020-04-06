package main

import (
	"fmt"
	"strconv"
)

// Person : struct for person
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + ", I am from " + p.city + "and, I am " + strconv.Itoa(p.age)
}

// increseAge method pointer receiver
func (p *Person) increseAge() {
	p.age++
}

// change lastName when female gets married
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Isaiah", lastName: "Yemitan", city: "lagos", gender: "m", age: 25}
	// alternative
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.increseAge()
	person1.getMarried("gh")
	fmt.Println(person1)

	person2 := Person{"laura", "Sar", "london", "f", 33}
	fmt.Println(person2)
	person2.getMarried("Yemitan")
	fmt.Println(person2)
}
