package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
The stdin of the Go process could be retrieved via the Stdin
of the os package. In fact, it is a File type which implements
the Reader interface. Reading from the Reader is then very easy.
The preceding code shows three very common ways of how to read
from the Stdin.

The scanner, which is the second option provides a convenient
way of scanning larger input. The Scanner contains the function
Split by which the custom split function could be defined.
For example, to scan the words from stdin, you can use
bufio.ScanWords predefined SplitFunc.
*/

func main() {
	// The Scanner is able to
	// scan input by lines
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
