package main

import "fmt"

var x int = 42
var y string = "James Bomd"
var z bool = true

func main(){
	s := fmt.Sprint(x,y,z)
	fmt.Println(s)
}
