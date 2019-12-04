package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying to people", func(t *testing.T) {
		got := Hello("Ahsan!", "English")
		want := "Hello, Ahsan!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying default to people", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying spanish to people", func(t *testing.T) {
		got := Hello("Ahsan!", "Spanish")
		want := "Hola, Ahsan!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying french to people", func(t *testing.T) {
		got := Hello("Ahsan!", "French")
		want := "Bonjour, Ahsan!"
		assertCorrectMessage(t, got, want)
	})
}
