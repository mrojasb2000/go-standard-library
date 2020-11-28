package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

/*
The golang.org/x/text/encoding/charmap package constains the Charmap
type pointer constants the represent the widely used charsets.
The Charmap type provides the methods for creating the encoder and
decoder for the given charset. The Encoder creates the encoding Writer
which encodes the written bytes into the chosen charset. Similarly,
the Decoder can create the decoding Reader, which decodes all read data
from the chosen charset.
test
*/
func main() {
	// Write the string
	// encoded to Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String("This is sample text with runes Š")
	if e != nil {
		panic(e)
	}
	ioutil.WriteFile("temp/example.txt", []byte(s), os.ModePerm)

	// Decoder to UTF-8
	f, e := os.Open("temp/example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
