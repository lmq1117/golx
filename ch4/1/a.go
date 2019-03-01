package main

import "fmt"

func main() {
	//var a [3]int
	var s [4]string
	//fmt.Println(a[0])
	//fmt.Println(len(a) - 1)

	for i, v := range s {
		//fmt.Printf("%d %d\n",i,v)
		fmt.Printf("%d --|%s|--\n", i, v)
	}
}
