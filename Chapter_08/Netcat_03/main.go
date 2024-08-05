// main.go

// Netcat_03 is a simple read/write client for TCP servers.

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors.
		log.Println("done")
		done <- struct{}{} // Signal the main goroutine.
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // Wait for background goroutine to finish.
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
