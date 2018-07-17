package goroutines

import (
	"testing"
)

var testCases = []struct {
	input    string
	expected string
}{
	{
		input:    "",
		expected: "()",
	},
	{
		input:    "Hello World",
		expected: "(Hello World)",
	},
	{
		input:    "tag 2.0.68",
		expected: "(tag 2.0.68)",
	},
	{
		input:    "Hey !! How are you?",
		expected: "(Hey !! How are you?)",
	},
	{
		input:    "(Hello World!!)",
		expected: "((Hello World!!))",
	},
	{
		input:    "12345",
		expected: "(12345)",
	},
}

func TestProcess(t *testing.T) {
	for _, tt := range testCases {
		// GIVEN
		input := make(chan string)
		defer close(input)

		done := make(chan bool)
		defer close(done)

		go func() {
			input <- tt.input
			done <- true
		}()

		// WHEN
		output := Process(input)
		<-done // blocks until the input write routine is finished

		// THEN
		found := <-output // blocks until the output has contents

		if found != tt.expected {
			t.Errorf("Expected %s, found %s", tt.expected, found)
		}
	}
}
