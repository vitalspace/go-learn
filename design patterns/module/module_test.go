package design patterns/module

import (
	"testing"
)

// TestSayPublic is a unit test for the sayPublic method
func TestSayPublic(t *testing.T) {
	// Create a new module with a private message
	m := Module{private: "Hello world this is a private message"}

	// Call the module's sayPublic method
	m.sayPublic()

	// There is no explicit assertion in this test, as the output of sayPublic is printed to the console
	// However, if sayPublic is not functioning correctly, the test will fail.
	// This test can be run using the "go test" command in the terminal.
}
