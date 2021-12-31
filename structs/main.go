package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	arnold := person{
		firstName: "Arnold",
		lastName:  "Babasa",
		contactInfo: contactInfo{
			email:   "realnsleo@gmail.com",
			zipCode: 256,
		},
	}

	// p := &arnold
	// fmt.Printf("%p", p)

	arnold.updateName("John")
	arnold.print()
}

func (p *person) updateName(n string) {
	p.firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
