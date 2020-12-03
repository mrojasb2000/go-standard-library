package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

/*
The io.Pipe function creates the in-memory pipe and returns both ends of the pipe,
the PipeReader on one side and PipeWriter on the other side. Each Write to PipeWriter
is blocked until it is consumed by Read on the other end.

The examplw shows the piping output from the excuted command to the standard output
of the parent program. By assigning the pWriter to cmd.Stdout, the standard output
of the child process is written to the pipe, and the io.Copy in goroutine consumes
the written data, by copying the data to os.Stdout.
*/
func main() {
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "Hello Go!\nThis is example")
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		if _, err := io.Copy(os.Stdout, pReader); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
