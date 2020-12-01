package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
The binary data could be written with the use of the encoding/binary package.
The function Write consumes the Writer where the data should be written, the
byte order (BigEndian/LittleEndian) and finally, the value to be writtem into
Writer.

To read binary data analogically, the read function could be used. Note that
there is no magic in reading the data from the binary source. You need to be
sure what data you are fetching from the Reader. If not, the data could be
fetching into any type which fits the size.
*/
func main() {
	// Writing binary values
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}
	if err := binary.Write(buf, binary.BigEndian, []byte("Hello")); err != nil {
		panic(err)
	}

	// Reading the written values
	var num float64
	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)
	greeting := make([]byte, 5)
	if err := binary.Read(buf, binary.BigEndian, &greeting); err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", greeting)
}
