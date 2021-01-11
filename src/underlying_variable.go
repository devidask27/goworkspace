package main

import "fmt"

type newType int

var x newType

func main() {
	fmt.Println(x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T", x)
}
