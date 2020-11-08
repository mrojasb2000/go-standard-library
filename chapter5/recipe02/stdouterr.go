package main

import (
	"fmt"
	"io"
	"os"
)

/*
As with the Stdin from the previos recipe, the Stdout and
Stderr are the file descriptors. So the are implementing
the Writer interface.

The preceding example shows a few ways of how to write into
these via the io.WriteString function, with the use of the Writer
API and by the fmt package and Fprintln funcitons.
*/
func main() {
	// Simply write string
	io.WriteString(os.Stdout, "This is string to standard output.\n")

	io.WriteString(os.Stderr, "This is string to standard error output.\n")

	// Stdout/err implements
	// writer interface
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// The fmt package
	// could be used too
	fmt.Fprintln(os.Stdout, "\n")

}
