package main

import "fmt"

func main(){
	// a SLICE allows you to group together VALUES of the same TYPE
	s := make([]int ,4)
	x := []int {4,5,7,8,12}
	//fmt.Println(s)
	//fmt.Println(x)
	for i,v := range x{
		fmt.Println(i,v)
	}
}
