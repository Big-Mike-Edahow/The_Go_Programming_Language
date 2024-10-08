// main.go

// Countdown_02 implements the countdown for a rocket launch.

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Create abort channel.
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read a single byte.
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
