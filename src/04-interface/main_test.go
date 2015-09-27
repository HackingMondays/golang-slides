package main

import "testing"

func TestHello(t *testing.T) {
	testCase(t, &EnglishSpeaker{}, "Hello")
	testCase(t, &GermanSpeaker{"Morgen"}, "Guten Morgen")
	testCase(t, &GermanSpeaker{"Abend"}, "Guten Abend")
}

func testCase(t *testing.T, speaker Speaker, expected string) {
	// call function Hello() on interface Speaker
	actual := speaker.Hello()

	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}
