package main

import (
	"strings"
	"fmt"
)

func main() {
	helloWorldArray := strings.Split("Hello World", "")

	fmt.Printf("Original Value: %v\n", helloWorldArray)
	fmt.Printf("As String     : %v\n", strings.Join(helloWorldArray, ""))

	fmt.Println("\nCounting For-Loop:")
	for i := 0; i < len(helloWorldArray); i++ {
		fmt.Printf("char (%2d): %v\n", i, helloWorldArray[i])
	}

	fmt.Println("\nRange For-Loop:")
	for i, char := range helloWorldArray {
		fmt.Printf("char (%2d): %v\n", i, char)
	}

	fmt.Println("\nSlice Actions:")
	fmt.Printf("Char 0 to 5         : %v\n", helloWorldArray[:5])
	fmt.Printf("Char 6 and following: %v\n", helloWorldArray[6:])
}
