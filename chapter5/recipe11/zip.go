package main

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
The built-in package zip contains the NewWriter and NewReader
functions to create the zip.Writer to compress, and the zip.Reader
to descompress the zipped content.

Each record of the ZIP file is created with the Create method of the
created zip.Writer. The returned Writer is then used to write the
content body.

To descompress the files, the OpenReader function is used to create the
ReadCloser of the records in the zipped file. The file field of the
created ReaderCloser is the slice of zip.file pointers ([]*zip.file).
The content of the file is obtained by calling the Open method and by
reading the returned ReadCloser.
*/
func main() {
	var buff bytes.Buffer

	// Compress content
	zipW := zip.NewWriter(&buff)
	f, err := zipW.Create("temp/newfile.txt")
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte("This is my file content"))
	if err != nil {
		panic(err)
	}

	err = zipW.Close()
	if err != nil {
		panic(err)
	}

	// Write output to file
	err = ioutil.WriteFile("temp/data.zip", buff.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Decompress the comtent
	zipR, err := zip.OpenReader("temp/data.zip")
	if err != nil {
		panic(err)
	}

	for _, file := range zipR.File {
		println("File ", file.Name+" contains:")
		r, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			panic(err)
		}
		err = r.Close()
		if err != nil {
			panic(err)
		}
		println()

	}

}
