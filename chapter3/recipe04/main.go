package main

import (
	"fmt"
	"math/big"
)

// PI const value
const PI = `3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196`
const diameter = 3.0
const precision = 400

func main() {
	pi, _ := new(big.Float).SetPrec(precision).SetString(PI)
	d := new(big.Float).SetPrec(precision).SetFloat64(diameter)

	circumference := new(big.Float).Mul(pi, d)

	pi64, _ := pi.Float64()
	fmt.Printf("Circunference big.Float = %.400f\n", circumference)

	fmt.Printf("Circunference float64 = %.400f\n", pi64*diameter)

	sum := new(big.Float).Add(pi, pi)
	fmt.Printf("Sum = %.400f\n", sum)

	diff := new(big.Float).Sub(pi, pi)
	fmt.Printf("Diff = %.400f\n", diff)

	quo := new(big.Float).Quo(pi, pi)
	fmt.Printf("Quo = %.400f\n", quo)

}
