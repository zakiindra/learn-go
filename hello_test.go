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
		actual := Hello("Dude")
		expected := "Hello, Dude"

		AssertCorrect(t, actual, expected)
	})

	t.Run("return Hello, world! if no argument supplied", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, world"

		AssertCorrect(t, actual, expected)
	})
}
