package main

import "fmt"

type Sender interface {
	Send(message string) error
}

type StdOut struct{}

func NewStdOut() StdOut {
	return StdOut{}
}

func (s StdOut) Send(message string) error {
	fmt.Println(message)

	return nil
}

func SendInBatch(sender []Sender, message string) error {
	var err error

	for _, s := range sender {
		if e := s.Send(message); e != nil {
			err = e
		}
	}

	return err
}

func main() {
	sender := []Sender{NewStdOut()}

	if err := SendInBatch(sender, "test message"); err != nil {
		panic(err)
	}
}
