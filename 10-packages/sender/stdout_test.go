package sender

import "testing"

func TestStdOut_Send(t *testing.T) {
	t.Run("should not fail", func(t *testing.T) {
		s := NewStdOut()

		if err := s.Send("test message"); err != nil {
			t.Errorf("expected not to fail, %v", err.Error())
		}
	})
}
