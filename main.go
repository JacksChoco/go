package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	fristName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		fristName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "zkmm234@naver.com",
			zipCode: 123,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).fristName = newFirstName
}
