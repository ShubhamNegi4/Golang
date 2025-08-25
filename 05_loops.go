package main

import (
	"fmt"
)

func main() {

	//for loop
	fmt.Println("For loop")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//there is no while loop, but we can use for loop like while loop
	fmt.Println("while loop")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//short hand
	fmt.Println("short hand")
	var str = []int{1, 2, 3, 4, 5, 6}
	for idx, val := range str {
		fmt.Println(idx, val)
	}

	//without idx
	fmt.Println()
	for _, val := range str {
		fmt.Println(val)
	}

}
