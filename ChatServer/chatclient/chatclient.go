package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// connects to a server address on a named network
	conn, err := net.Dial("tcp", "localhost:9090")

	if err != nil {
		// raised when an unexpected error occurs
		panic(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
