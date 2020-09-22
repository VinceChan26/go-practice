package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go run on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS", runtime.GOARCH)
	default:
		fmt.Println(os, runtime.GOARCH)
	}
}
