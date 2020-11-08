package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

/*
The reading from the file is simple because the File type
implements both the Reader and Writer interfaces. This way,
all functions and approaches applicable to the Reader interface
are applicable to the File type. The precending example shows to
read the file with the use of Scanner and write the content to the
bytes buffer (which is more performant than string concatenation).
This way, you are able to control the volume of content read from
file.

The second approach with ioutil.ReadFile is simpler but should be
used carefully, because it reads the whole file. Keep in mind
that the file could be huge and it could threaten the stabily of
the application.
*/
func main() {
	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the
	// file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// for smaller files
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
