package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaaaa"
	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}
