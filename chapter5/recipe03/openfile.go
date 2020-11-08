package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
The os package offers aa simple way of opening the file.
The function Open opens the file by the path, just in read-only mode.
Another function, OpenFile, is the more powerful one and consumes the
path to the file, flags, and permissions.

The flag constants are defined in the os package and you can
combine them with use of the binary OR operator |. The permissions
are set by the os package constantes (for example, os.ModePerm)
or by the number notation such as 0777 (permissions: rwxrwxrwx).
*/
func main() {
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile("temp/test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test String")
	f.Close()
}
