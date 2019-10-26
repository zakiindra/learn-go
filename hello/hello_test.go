package main

import "testing"

func TestHello(t *testing.T) {

	AssertCorrect := func(t *testing.T, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("got %q want %q", actual, expected)
		}
	}

	t.Run("Should return greeting with name", func(t *testing.T) {
		actual := Hello("Dude", "English")
		expected := "Hello, Dude"

		AssertCorrect(t, actual, expected)
	})

	t.Run("Return Hello, world! if no argument supplied", func(t *testing.T) {
		actual := Hello("", "English")
		expected := "Hello, world"

		AssertCorrect(t, actual, expected)
	})

	t.Run("Should greet in Indonesian with name", func(t *testing.T) {
		actual := Hello("Nia", "Indonesian")
		expected := "Halo, Nia"

		AssertCorrect(t, actual, expected)
	})

	t.Run("Should greet in French with name", func(t *testing.T) {
		actual := Hello("Nia", "French")
		expected := "Bonjour, Nia"

		AssertCorrect(t, actual, expected)
	})

}
