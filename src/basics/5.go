package main

import "fmt"

var c, golang bool

func main() {
	var i int
	var nilChannel chan int
	channel := make(chan int)
	fmt.Print(i, c, golang, nilChannel, channel)
}

