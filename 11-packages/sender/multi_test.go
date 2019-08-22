package sender

import (
	"errors"
	"testing"
)

type FakeSender struct {
	ShouldFail bool
	messages   []string
}

func NewFakeSender(shouldFail bool) FakeSender {
	return FakeSender{
		ShouldFail: shouldFail,
	}
}

func (f *FakeSender) Send(message string) error {
	if f.ShouldFail {
		return errors.New("should fail")
	}

	f.messages = append(f.messages, message)

	return nil
}

func TestMulti_Send(t *testing.T) {
	t.Run("should forward message to all senders", func(t *testing.T) {
		fs1 := NewFakeSender(false)
		fs2 := NewFakeSender(false)

		s := NewMulti([]Sender{&fs1, &fs2})

		if err := s.Send("test message"); err != nil {
			t.Errorf("expected not to fail, %v", err.Error())
		}

		if len(fs1.messages) != 1 {
			t.Error("expected sender 1 to have one message")
		}

		if fs1.messages[0] != "test message" {
			t.Error("expected sender 1 to get an untouched message")
		}

		if len(fs2.messages) != 1 {
			t.Error("expected sender 2 to have one message")
		}

		if fs2.messages[0] != "test message" {
			t.Error("expected sender 2 to get an untouched message")
		}
	})

	t.Run("should fail to send message to one sender", func(t *testing.T) {
		fs1 := NewFakeSender(false)
		fs2 := NewFakeSender(true)

		s := NewMulti([]Sender{&fs1, &fs2})

		if err := s.Send("test message"); err == nil {
			t.Error("expected sending to fail")
		}

		if len(fs1.messages) != 1 {
			t.Error("expected sender 1 to have one message")
		}

		if fs1.messages[0] != "test message" {
			t.Error("expected sender 1 to get an untouched message")
		}

		if len(fs2.messages) != 0 {
			t.Error("expected sender 2 to have no message")
		}
	})
}
