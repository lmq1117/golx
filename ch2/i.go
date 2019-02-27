package main

import "fmt"

func main(){
    var p = f()
    var q = f()
    fmt.Println(p)
    fmt.Println(q)
}

func f() * int {
    v := 1
    return &v
}

