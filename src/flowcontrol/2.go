package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		if x := rand.Intn(30); x < 10 {
			fmt.Println("x < 10")
		} else if x > 10 && x < 20 {
			fmt.Println("x > 11")
		} else {
			fmt.Println("x > 20")
		}
	}
}
