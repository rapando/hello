package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(AddNumbers(4, 5))
	fmt.Println(IsGreaterThan5(6)
}

func AddNumbers(x, y int) int {
	return x + y
}

func IsGreaterThan5(x int) bool {
	return x > 5
}
