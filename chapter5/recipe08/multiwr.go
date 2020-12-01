package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
The io package contains the MultiWriter function with variadic parameters of Writers.
When the Write method on the Writer is called, then the data is written to all
underlying Writers.
*/
func main() {
	buf := bytes.NewBuffer([]byte{})
	f, err := os.OpenFile("temp/sample.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	wr := io.MultiWriter(buf, f)
	_, err = io.WriteString(wr, "Hello, Go is awesome!")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of buffer: " + buf.String())
}
