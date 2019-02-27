package main

import "fmt"

func main() {
	v := 1
	fmt.Println(v)
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	//通过传地址改变v的值
	*p++
	return *p
}
