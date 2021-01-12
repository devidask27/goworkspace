package main

import "fmt"

func main(){
	//The capacity of a slice is the number of elements for which there is
	// space allocated in the underlying array. At any time the following
	// relationship holds 0 <= len(s) <= cap(s)

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x, 3423)
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x, 3424)
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

