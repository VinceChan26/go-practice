package main

import (
	"fmt"
	time2 "time"
)

func main() {
	time := time2.Now()
	fmt.Println("Now hour:", time.Hour())
	switch {
	case time.Hour() < 12:
		fmt.Println("Good morning")
	case time.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
