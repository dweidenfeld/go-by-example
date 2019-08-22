package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("You have to enter a number")
	}

	firstNumber, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	if firstNumber > 0 {
		fmt.Println("Your number is bigger than 0")
	}

	if firstNumber < 5 {
		fmt.Println("Your number is lower than 5")
	}

	switch firstNumber {
	case 0, 1:
		fmt.Println("Your number is 0 or 1")
	case 2, 3:
		fmt.Println("Your number is 2 or 3")
	case 4:
		fmt.Println("You chose 4")
	default:
		fmt.Println("You case is not handled in this switch statement")
	}
}
