package podcast

import "testing"

func TestITunesID(t *testing.T) {
	t.Run("with id prefix", func(t *testing.T) {
		id := ITunesID("id1234567890")
		i, err := id.Int()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if i != 1234567890 {
			t.Errorf("expected 1234567890, got %d", i)
		}
	})

	t.Run("with numeric string only", func(t *testing.T) {
		id := ITunesID("1234567890")
		i, err := id.Int()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if i != 1234567890 {
			t.Errorf("expected 1234567890, got %d", i)
		}
	})

	t.Run("with invalid format", func(t *testing.T) {
		id := ITunesID("asdasdba")
		i, err := id.Int()
		if err == nil {
			t.Errorf("expected error, got %d", i)
		}
	})
}
