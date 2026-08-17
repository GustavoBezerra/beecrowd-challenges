package main

import "testing"

func Test1000(t *testing.T) {
	expected := "Hello World!"
	actual := helloWorld()

	if actual != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, actual)
	}
}
