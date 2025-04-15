package podcast

import "testing"

func TestITunesID(t *testing.T) {
	id := ITunesID("id1234567890")
	if id.Int() != 1234567890 {
		t.Errorf("expected 1234567890, got %d", id.Int())
	}
	id = ITunesID("1234567890")
	if id.Int() != 1234567890 {
		t.Errorf("expected 1234567890, got %d", id.Int())
	}
	id = ITunesID("asdasdba")
	if id.Int() != 0 {
		t.Errorf("expected 0, got %d", id.Int())
	}
}
