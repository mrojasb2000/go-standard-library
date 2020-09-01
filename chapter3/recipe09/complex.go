package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// complex numbers are
	// defined as real and imaginary
	// part defined by float64
	a := complex(2, 3)

	fmt.Printf("Real part: %f \n", real(a))
	fmt.Printf("Complex par: %f \n", imag(a))

	b := complex(6, 4)

	// All common
	// operators are useful
	c := a - b
	fmt.Printf("Difference : %v\n", c)

	c = a + b
	fmt.Printf("Sum : %v\n", c)

	c = a * b
	fmt.Printf("Product : %v\n", c)

	c = a / b
	fmt.Printf("Division : %v\n", c)

	conjugate := cmplx.Conj(a)
	fmt.Printf("Complex number a's conjugate : %v\n", conjugate)

	cos := cmplx.Cos(b)
	fmt.Printf("Cosine of b : %v\n", cos)

}
