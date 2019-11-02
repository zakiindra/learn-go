package main

import (
    "bytes"
    "testing"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Nia")

    actual := buffer.String()
    expected := "Hello, Nia"

    if actual != expected {
        t.Errorf("Expected %q but got actual %q", expected, actual)
    }
}
