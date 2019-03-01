package main

import (
	"fmt"
	"math"
)

func main() {
	consts()
	enums()
}

func consts() {
	const filename string = "file name const"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, a, b, c)
}

func enums() {
	const (
		java = 0
		cpp  = 1
		c    = 2
		php  = 3
	)
	fmt.Println(java, cpp, c, php)

	// 使用iota实现自增枚举
	const (
		man = iota
		women
		shemale
	)
	fmt.Println(man, women, shemale)
}
