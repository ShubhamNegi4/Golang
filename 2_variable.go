//note- every imported package and declared variable must be used; otherwise you will get a compile time error
// for variables you can declear any variable with _ if you don't want to use it and it won't show any compile error anymore

package main

import "fmt"

var sample0 int //global

func main() {
	var sample = 45              // auto detect the datatype and store the value
	var sample2 string = "hello" // we can manually set the datatype
	var sample3 float32          // if we don't initilize a varible, we get the default value as 0
	fmt.Print(sample, sample2)   // output 45hello
	fmt.Println()                // for extra endline
	fmt.Println(sample, sample2) //output -	45 hello

	fmt.Println(sample3) // output - 0

	sample3 = 1 // we can assign any value after the declaration and update it's value
	fmt.Println(sample3)

	//short hand of line 7 is
	sample4 := 45        //but we can't use this outside the function(globally)
	fmt.Println(sample4) // output -> 45

	fmt.Println(sample0) // output -> 0

	fmt.Printf("%s\n", sample2)

	//Sprint (save formatted string)
	var str = fmt.Sprintln(sample2)
	fmt.Print(str)
}

/* output of above code

shubhamnegi@senkus-MacBook-Air Golang % go run  2_variable.go
45hello
45 hello
0
1
45
0
hello
hello
shubhamnegi@senkus-MacBook-Air Golang %

*/
