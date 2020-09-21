package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println( "result:", sum)

	// while
	var count float64 = 0
	var base float64
	for count < 5 {
		base = math.Pow(2, count)
		count++
	}
	fmt.Print("result: ", base)
}
