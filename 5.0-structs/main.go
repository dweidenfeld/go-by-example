package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(firstName, lastName string, age int) (Person, error) {
	if firstName == "" || lastName == "" {
		return Person{}, errors.New("names must not be empty")
	}

	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}, nil
}

func main() {
	dominik := Person{
		FirstName: "Dominik",
		LastName:  "Weidenfeld",
		Age:       29,
	}
	fmt.Printf("Dominik             : %v\n", dominik)
	fmt.Printf("Dominik (formatted) : %+v\n", dominik)

	fmt.Println()

	rainer, err := NewPerson("Rainer", "", 0)
	if err != nil {
		fmt.Printf("Error creating Rainer: %v\n", err)
	}
	fmt.Printf("Rainer              : %v\n", rainer)
	fmt.Printf("Rainer (formatted)  : %+v\n", rainer)
}
