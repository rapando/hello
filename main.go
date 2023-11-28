package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(AddNumbers(4, 5))
}

func AddNumbers(x, y int) int {
	return x + y
}
