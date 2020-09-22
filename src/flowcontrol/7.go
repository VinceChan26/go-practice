package main

import "fmt"

func main() {
	// stacking defers is LIFO (last-in-first-out)
	// 2 4 6 7 5 3 1
	defer fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
	defer fmt.Println("5")
	R()
	T()
}

func T() {
	defer fmt.Println("7")
}

func R() {
	fmt.Println("6")
}
