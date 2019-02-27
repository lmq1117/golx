package main

import "fmt"

func main() {
	var one = newIntOne()
	var two = newIntTwo()
	fmt.Println(*one)
	fmt.Println(one)

	fmt.Println(*two)
	fmt.Println(two)
}

func newIntOne() *int {
	return new(int)
}

func newIntTwo() *int {
	var dummy int
	return &dummy
}
