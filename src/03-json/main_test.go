package main

import "testing"

const jsonAsString = `{"EN":"Hello", "DE":"Guten Tag"}`

func TestHello(t *testing.T) {

	// Expect Hello() to return the German greeting
	expected := "Guten Tag"
	actual := Hello(jsonAsString)

	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}
