package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Person struct {
	FirstName string
	LastName  string
}

func NewPerson(firstName, lastName string) (Person, error) {
	if firstName == "" || lastName == "" {
		return Person{}, errors.New("names must not be empty")
	}

	return Person{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func (p Person) FullName() string {
	return fmt.Sprintf("%v %v", p.FirstName, p.LastName)
}

func main() {
	dominik, err := NewPerson("Dominik", "Weidenfeld")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Dominik: %v\n", dominik.FullName())
}
