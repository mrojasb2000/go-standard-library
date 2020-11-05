package main

import (
	"fmt"
	"os"
)

/*
The stdin of the Go process could be retrieved via the Stdin
of the os package. In fact, it is a File type which implements
the Reader interface. Reading from the Reader is then very easy.
The preceding code shows three very common ways of how to read
from the Stdin.

The reading via the Reader API is the last presented approach.
This one provides you with mora control of how the imput is read.
*/

func main() {
	for {
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}
}

func process(data []byte) {
	fmt.Printf("Received: %X %s\n", data, string(data))
}
