package main

import (
	"fmt"
	"sync"
)

// message is a global string that multiple goroutines will try to update.
// Access to this variable must be synchronized to avoid data races.
var message string

// waitgroup is used to wait for all launched goroutines to finish before
// allowing main to exit.
var waitgroup sync.WaitGroup

// updateMessage updates the global `message` variable. It receives a pointer
// to a sync.Mutex which it uses to protect the critical section (the write
// to `message`). The function calls waitgroup.Done() when it completes so
// that the caller can wait for all updates.
func updateMessage(newMessage string, mutex *sync.Mutex) {
	// Ensure the WaitGroup is decremented when this goroutine returns.
	defer waitgroup.Done()

	// Lock the mutex before modifying the shared variable. Use defer to
	// guarantee the mutex is released even if the function grows or returns
	// early in the future.
	mutex.Lock()
	defer mutex.Unlock()

	// Critical section: write to the shared variable.
	message = newMessage
}

func main() {

	message = "Hello, World!"

	// Initialize the shared state.
	message = "Hello, World!"

	// mutex protects concurrent access to `message`.
	var mutex sync.Mutex

	// We will start 2 goroutines that update `message`.
	waitgroup.Add(2)

	// Launch goroutines, passing a pointer to the mutex so each goroutine
	// can lock before writing.
	go updateMessage("Hello, Go!", &mutex)
	go updateMessage("Hello, Concurrency!", &mutex)

	// Wait for both goroutines to finish.
	waitgroup.Wait()

	// Print the final value of message. Because updates are synchronized,
	// there is no data race here.
	fmt.Println(message)
}
