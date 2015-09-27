package main

import "net/http"
import "net/http/httptest"
import "testing"

func TestHello(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://dummy", nil)
	myHandler(recorder, request)

	expected := "Hello World"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}
