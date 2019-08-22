package main

import (
	"fmt"
)

func main() {
	chars := make(chan string)

	go func() {
		defer close(chars)
		for i := 0; i < 10; i++ {
			chars <- "A"
		}
	}()

	for c := range chars {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
