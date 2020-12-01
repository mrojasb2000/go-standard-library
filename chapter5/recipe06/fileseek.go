package main

import (
	"errors"
	"fmt"
	"os"
)

/*
The preceding example uses the flatfile as an illustration of how to seek
, read and write at the position in the file. In general, for moving the
position of the current pointer in the File, the Seek method can be used.
It takes two arguments and these are, position and how to count the
position,
	0 - relative to file origin,
	1 - relative to current position,
	2 - relative to the end of file.
This way you are able to move the cursor within the file. The Seek method
is used in the implementation of the readLine function in the precending
code.

flatfile:
123.Jun.......Wong......
12..Novak.....Jurgen....
10..Thomas....Sohlich...
*/
const lineLength = 25

func main() {
	f, e := os.OpenFile("temp/flatfile.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	fmt.Println(readRecords(2, "last", f))
	if err := writeRecord(2, "first", "Radomir", f); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(2, "first", f))
	if err := writeRecord(10, "first", "Andrew", f); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(10, "first", f))
	fmt.Println(readLine(2, f))
}

func readLine(line int, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.Seek(int64(line*lineLength), 0)
	_, err := f.Read(lineBuffer)
	return string(lineBuffer), err
}

func writeRecord(line int, column, dataStr string, f *os.File) error {
	definedLen := 10
	position := int64(line * lineLength)
	switch column {
	case "id":
		definedLen = 4
	case "first":
		position += 4
	case "last":
		position += 14
	default:
		return errors.New("Column not defined")
	}

	if len([]byte(dataStr)) > definedLen {
		return fmt.Errorf("Maximum length for '%s' is %d", column, definedLen)
	}

	data := make([]byte, definedLen)
	for i := range data {
		data[i] = '.'
	}
	copy(data, []byte(dataStr))
	_, err := f.WriteAt(data, position)
	return err
}

func readRecords(line int, column string, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.ReadAt(lineBuffer, int64(line*lineLength))
	var retVal string
	switch column {
	case "id":
		return string(lineBuffer[:3]), nil
	case "first":
		return string(lineBuffer[4:13]), nil
	case "last":
		return string(lineBuffer[14:23]), nil
	}
	return retVal, errors.New("Column not defined")
}
