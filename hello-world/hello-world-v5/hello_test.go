package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ahsan!")
		want := "Hello, Ahsan!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello, world! for empty name", func(t *testing.T){
		got := Hello("")
		want := "Hello, world!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
