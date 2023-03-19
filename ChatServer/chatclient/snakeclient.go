package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
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
	inputs := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b []byte = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			ch <- string(b)
		}
	}(inputs)

	for {
		select {
		case stdin, _ := <-inputs:
			fmt.Println("Keys pressed:", stdin)
			mustCopy(conn, stdin)
		default:
			// fmt.Println("Working..")
		}
		time.Sleep(time.Millisecond * 100)
	}

	// conn.Close()
	// <-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, msg string) {
	if _, err := fmt.Fprintln(dst, msg); err != nil {
		log.Fatal(err)
	}
}
