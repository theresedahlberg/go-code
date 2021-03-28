package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Initalize person using struct
	person1 := Person{firstName: "Therese", lastName: "Dahlberg", city: "Göteborg", gender: "f", age: 22}
	//fmt.Println(person1)

	//Alternative way
	//person2 := Person{"Therese", "Dahlberg", "Göteborg", "f", 22}

	//fmt.Println(person2)
	//fmt.Println(person1.firstName)

	//person1.firstName = "danthe"
	//fmt.Println(person1.firstName)
	//person1.age++
	//fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())

}
