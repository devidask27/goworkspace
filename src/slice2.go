package main

import "fmt"

func main(){
	s := []int{1,2,3}
	fmt.Println("slice is : ",s)
	s = append(s,4,5,6)
	fmt.Println("slice after appended ",s)
}
