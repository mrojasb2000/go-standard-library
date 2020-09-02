package main

import (
	"fmt"
	"math"
)

// Radian type
type Radian float64

// ToDegrees return Degress instance type
func (rad Radian) ToDegrees() Degree {
	return Degree(float64(rad) * (180.0 / math.Pi))
}

// Float64 return instance type
func (rad Radian) Float64() float64 {
	return float64(rad)
}

// Degree type
type Degree float64

// ToRadians return Radian instance type
func (deg Degree) ToRadians() Radian {
	return Radian(float64(deg) * (math.Pi / 180.0))
}

// Float64 return instance type
func (deg Degree) Float64() float64 {
	return float64(deg)
}

func main() {
	val := radiansToDegrees(1)
	fmt.Printf("One radian is : %.4f degrees\n", val)

	val2 := degreeToRadians(val)
	fmt.Printf("%.4f degrees is %.4f rad\n", val, val2)

	// Conversion as part
	// of type methods
	val = Radian(1).ToDegrees().Float64()
	fmt.Printf("Degrees: %.4f degrees\n", val)

	val = Degree(val).ToRadians().Float64()
	fmt.Printf("Rad: %.4f radians\n", val)
}

func degreeToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func radiansToDegrees(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}
