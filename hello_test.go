package main

import "testing"

func TestHello(t *testing.T) {
	
	assertCorrectMessage := func (t *testing.T, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s'. Want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Michael", "")
		want := "Hello, Michael!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world!' when  an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Jakob", "Spanish")
		want := "Hola, Jakob!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("and french", func(t *testing.T) {
		got := Hello("Jakob", "French")
		want := "Bonjour, Jakob!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("and finally danish", func(t *testing.T) {
		got := Hello("Åge", "Danish")
		want := "Hejsa, Åge!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Default to english if the language is not implemented", func(t *testing.T) {
		got := Hello("Hansie", "SomeUnknownLanguage")
		want := "Hello, Hansie!"
		assertCorrectMessage(t, got, want)
	})
}