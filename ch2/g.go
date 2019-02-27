package main

import "fmt"

func main(){
    x := 1
    p := &x

    fmt.Println(x)
    fmt.Println(p)
    *p = 2
    fmt.Println(x)
    fmt.Println(p)
    x = 30
    fmt.Println(x)
    fmt.Println(p)

}

