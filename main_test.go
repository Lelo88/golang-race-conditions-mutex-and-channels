package main

import "testing"

func Test_updateMessage(t *testing.T) {

	// Use the package-level variable `message` (do not re-declare a local
	// variable). The test sets an initial value, launches a goroutine that
	// updates it, waits, and then checks the final value.
	message = "Initial Message"

	waitgroup.Add(1)
	go updateMessage("Is it a final message?")
	waitgroup.Wait()

	if message != "Is it a final message?" {
		t.Errorf("Expected message to be updated to 'Is it a final message?', but got '%s'", message)
	}
}
