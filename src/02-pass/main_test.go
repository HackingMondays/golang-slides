package main

import "testing"

func TestHello(t *testing.T) {

	// Expect function Hello() to return "Hello"
	expected := "Hello"
	actual := Hello()

	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}
