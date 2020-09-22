package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		temp := z - (z*z-x)/(2*z)
		fmt.Println(temp)
		if temp == z || math.Abs(temp-z) < 0.000000000001 {
			break
		}
		z = temp
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
