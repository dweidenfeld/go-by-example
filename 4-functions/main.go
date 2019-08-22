package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	result1 := HelloWorld()
	fmt.Printf("HelloWorld()\t\t\t\t\t: %v\n", result1)

	result2, err := HelloWorldWithName("")
	fmt.Printf("HelloWorldWithName(\"%v\")\t\t\t\t: \"%v\" : %v\n", "", result2, err)

	input3 := "Service Community"
	result3, err := HelloWorldWithName(input3)
	fmt.Printf("HelloWorldWithName(\"%v\")\t\t: \"%v\" : %v\n", input3, result3, err)

	input4 := "Service Community"
	result4, len4, err := HelloWorldWithNameAndLen(input4)
	fmt.Printf("HelloWorldWithNameAndLen(\"%v\")\t: \"%v\" (%v) : %v", input4, result4, len4, err)
}

func HelloWorld() string {
	return "Hello World"
}

func HelloWorldWithName(name string) (string, error) {
	if name == "" {
		return "", errors.New("cannot handle empty input")
	}

	return fmt.Sprintf("Hello %v", name), nil
}

func HelloWorldWithNameAndLen(name string) (string, int, error) {
	formattedString, err := HelloWorldWithName(name)
	if err != nil {
		return "", -1, err
	}

	return formattedString, len(formattedString), nil
}
