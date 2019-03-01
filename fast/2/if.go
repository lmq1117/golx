package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile()
}

func readFile() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}
	fmt.Printf("%s", contents)
}
