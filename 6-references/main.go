package main

import "fmt"

type Message struct {
	Content string
}

func NewMessage(content string) Message {
	return Message{Content: content}
}

func NewMessageRef(content string) *Message {
	return &Message{Content: content}
}

func ManipulateMessage(message Message) {
	message.Content = "something different"
}

func ManipulateMessageRef(message *Message) {
	message.Content = "something different by ref"
}

func main() {
	msg1 := NewMessage("Message 1")
	ManipulateMessage(msg1)
	fmt.Printf("msg1: %v\n", msg1)

	msg2 := NewMessageRef("Message 2")
	ManipulateMessage(*msg2)
	fmt.Printf("msg2: %v\n", msg2)

	msg3 := NewMessage("Message 3")
	ManipulateMessageRef(&msg3)
	fmt.Printf("msg3: %v\n", msg3)

	msg4 := NewMessageRef("Message 4")
	ManipulateMessageRef(msg4)
	fmt.Printf("msg4: %v\n", msg4)
}
