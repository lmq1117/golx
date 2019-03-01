package main

import "fmt"

//计算两个数的最大公约数
func main() {
	//fmt.Println(gcd(8,12))
	test()
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func test() {
	x := 5
	y := 10

	x, y = y, x*x*y
	fmt.Println(x, y)
}
