package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * Random Number
 */
func main()  {
	i := 0
	for i < 5 {
		// change seeding
		rand.Seed(time.Now().UnixNano())
		// random 0 <= n < 10
		fmt.Printf("random number integer: %v\n", rand.Intn(10))
		i++
	}
}