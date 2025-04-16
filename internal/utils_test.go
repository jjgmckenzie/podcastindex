package internal

import (
	"testing"
)

func Test_BoolToInt(t *testing.T) {
	if BoolToInt(true) != 1 {
		t.Errorf("BoolToInt(true) = %d; want 1", BoolToInt(true))
	}
	if BoolToInt(false) != 0 {
		t.Errorf("BoolToInt(false) = %d; want 0", BoolToInt(false))
	}
}
