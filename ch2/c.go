package main

import (
    "fmt"
    "os"
)

func main(){
    var f, err = os.Open("a.go")
    fmt.Println(f)
    fmt.Println(err)
}

