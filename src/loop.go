package main

import "fmt"

func main() {
	var array [4]int
	for i := 0; i < 4; i++ {
		array[i] = i
	}
	for i := 0; i < 4; i++ {
		fmt.Println(array[i])
	}
}
