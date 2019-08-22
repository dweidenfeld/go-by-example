package sender

import "fmt"

type StdOut struct{}

func NewStdOut() StdOut {
	return StdOut{}
}

func (s StdOut) Send(message string) error {
	fmt.Println(message)

	return nil
}
