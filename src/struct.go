package main

import "fmt"

func main() {
	type info struct {
		name        string
		phoneNumber int
	}
	var myInfo info
	myInfo.name = "devi"
	myInfo.phoneNumber = 123
	fmt.Println(myInfo)
}
