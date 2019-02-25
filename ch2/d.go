package main

import "fmt"

func main(){
    i := 100
    j := 200
    fmt.Println(i)
    fmt.Println(j)
    i, j = j, i
    fmt.Println(i)
    fmt.Println(j)
    fmt.Println(j)

}

