package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	lookFor := "lamb"
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	startWith := "Mary"
	starts := strings.HasPrefix(refString, startWith)
	fmt.Printf("The \"%s\" starts with \"%s\": %t \n", refString, startWith, starts)

	endWith := "Mary"
	ends := strings.HasPrefix(refString, startWith)
	fmt.Printf("The \"%s\" ends with \"%s\": %t \n", refString, endWith, ends)
}
