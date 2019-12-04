package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say to people", func(t *testing.T) {
		got := Hello("Ahsan!", "")
		want := "Hello, Ahsan!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say default hello world to people", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say to people for Bengali", func(t *testing.T) {
		got := Hello("এহসান!", "Bengali")
		want := "হ্যালো, এহসান!"
		assertCorrectMessage(t, got, want)
	})

}
