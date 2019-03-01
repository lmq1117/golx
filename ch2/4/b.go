package main

import "fmt"

//元组赋值

func main() {
	var x, y int
	x = 2
	y = 3
	fmt.Println(x)
	fmt.Println(y)
	x, y = y, x
	fmt.Println(x)
	fmt.Println(y)

}
