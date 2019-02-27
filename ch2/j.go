package main

import "fmt"

func makePoint() *int {
	v := 1
	return &v
}

func main() {
	var p = makePoint()
	fmt.Println(p)
	fmt.Println(*p)
}
