package main

import "fmt"

func main() {
	x, y := 101, 1667
	pointer := &y // pointer 指向 y
	fmt.Println(x, y, *pointer)
	if y == *pointer {
		fmt.Println("y and pointer value is equal")
	}
	*pointer = 103 // 通過 pointer 設置 y
	fmt.Println(x, y, *pointer)
	pointer = &x // pointer 指向 x
	*pointer = *pointer / 11 // pointer 運算
	fmt.Println(x, y, *pointer, pointer) // x 與 pointer value 一致

}
