package main

import "fmt"

func main() {
	PrintSomething("Dominik")

	go PrintSomething("Rainer")
}

func PrintSomething(name string) {
	fmt.Printf("Hello %v\n", name)
}
