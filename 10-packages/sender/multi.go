package sender

type Multi struct {
	sender []Sender
}

func NewMulti(sender []Sender) Multi {
	return Multi{
		sender: sender,
	}
}

func (m Multi) Send(message string) error {
	var err error

	for _, s := range m.sender {
		if e := s.Send(message); e != nil {
			err = e
		}
	}

	return err
}
