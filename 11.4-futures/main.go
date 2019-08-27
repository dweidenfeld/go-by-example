package main

import (
	"fmt"
	"strings"
)

func main() {
	nameChannel, err := toName("dominik weidenfeld")
	if err != nil {
		panic(err)
	}

	formattedName := <-nameChannel
	fmt.Printf("Name: %v\n", formattedName)

	fmt.Println()

	nameChunkChannel, err := toChunks(formattedName)
	if err != nil {
		panic(err)
	}

	for chunk := range nameChunkChannel {
		fmt.Printf("Name chunk: %v\n", chunk)
	}
}

func toName(name string) (chan string, error) {
	retChan := make(chan string, 1)

	go func(n string) {
		defer close(retChan)

		retChan <- strings.Title(n)
	}(name)

	return retChan, nil
}

func toChunks(name string) (chan string, error) {
	retChan := make(chan string)

	go func(n string) {
		defer close(retChan)

		for _, chunk := range strings.Split(n, " ") {
			retChan <- chunk
		}
	}(name)

	return retChan, nil
}
