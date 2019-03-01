package main

import (
	"fmt"
	"math"
)

/**
内建变量类型
	bool
	string
	(u)int	(u)int8 (u)int16 (u)int32 (u)int64
	uintptr 指针
	byte
	rune字符型 32位 类比char
	float32 float64
	complex32 complex64 复数 i^2 = 1
*/

//定义包内变量
var (
	aa = 1
	bb = "tian ya"
)

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	triangle() //强制类型转换
}

//定义变量 只使用默认值
func variableZeroValue() {
	var a int
	var s string
	var b bool
	fmt.Println(a, b, s)
}

//定义变量 赋初值
func variableInitialValue() {
	var a, b int = 1, 2
	b = 3
	var s string = "haha"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

//最简单定义变量方式
func variableShorter() {
	a, b, c, d := 1, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

func triangle() {
	var a, b = 5, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	var d float64 = 11.7
	var e int = int(d) //文章上说类型转换可以不加小括号 其实是要加的 ?????
	fmt.Println(e)
	fmt.Println(c)
}
