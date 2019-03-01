package main

import "fmt"

func main() {
	var x int
	x = 1 //命名变量的赋值
	fmt.Println(x)
	var p = &x
	fmt.Println(p)
	*p = 2 //通过指针间接赋值
	fmt.Println(x)

	//结构体字段赋值 todo
	//数组 slice map元素赋值 todo
	// *=
	// i++ 语句 不是表达式

}
