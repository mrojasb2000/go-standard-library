package main

import "fmt"

/*
The stdin of the Go process could be retrieved via the Stdin
of the os package. In fact, it is a File type which implements
the Reader interface. Reading from the Reader is then very easy.
The preceding code shows three very common ways of how to read
from the Stdin.

It option illustrates the use of the fmt package, which provides
the functions Scan, and ScanIn. The Scanf function reads the input
into given variable(s). The advantage of Scanf is that you can
determine the format of the scanned value. The Scan function just
reads the input into a variable (without pre-defined formatting)
and ScanIn, as its name suggests, reads the input ended with the
line break.
*/
func main() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("What is you age?")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hello %s, your age is %d\n", name, age)
}
