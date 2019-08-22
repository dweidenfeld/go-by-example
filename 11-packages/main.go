package main

import "github.com/dweidenfeld/go-by-example/11-packages/sender"

func main() {
	s := sender.NewMulti([]sender.Sender{sender.NewStdOut()})

	if err := s.Send("test message"); err != nil {
		panic(err)
	}
}
