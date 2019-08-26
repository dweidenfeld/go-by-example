package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var someArray []string
	// mutex := new(sync.Mutex)

	go func() {
		// mutex.Lock()
		for i := 0; i < 10; i++ {
			someArray = append(someArray, "A")
		}
		// mutex.Unlock()
	}()

	go func() {
		// mutex.Lock()
		for i := 0; i < 10; i++ {
			someArray = append(someArray, "B")
		}
		// mutex.Unlock()
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Result: %v", strings.Join(someArray, ""))
}
