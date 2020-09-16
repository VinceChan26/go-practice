package main

import "fmt"

func main() {
	x, y := 3, 5
	fmt.Println(x, y)
	x, y = swap(x, y)
	fmt.Print(x, y)
}

func swap(x, y int) (int, int) {
	return y, x
}
