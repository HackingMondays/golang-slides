package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {
	// redirect output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	grep("^root", "passwd")

	// reset output again
	w.Close()
	os.Stdout = old

	captured, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("Error: ", err)
	}
	actual := string(captured)
	expected := "root:*:0:0:System Administrator:/var/root:/bin/sh\n"

	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}
