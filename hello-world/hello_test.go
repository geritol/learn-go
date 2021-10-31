package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("custom hello", func(t *testing.T) {
		got := Hello("Gergő", "english")
		want := "Hello, Gergő"

		assertCorrectMessage(t, got, want)
	})

	t.Run("default hello", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Abel", "french")
		want := "Bonjour, Abel"
		assertCorrectMessage(t, got, want)
	})
}
