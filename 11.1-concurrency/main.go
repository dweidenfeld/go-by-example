package main

import (
	"fmt"
	"time"
)

func main() {
	PrintSomething("Dominik")

	go PrintSomething("Rainer")

	time.Sleep(1 * time.Millisecond)
}

func PrintSomething(name string) {
	fmt.Printf("Hello %v\n", name)
}
