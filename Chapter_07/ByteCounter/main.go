// main.go

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.

package main

import (
	"fmt"
)

type ByteCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // Reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Convert int to ByteCounter.
	return len(p), nil
}
