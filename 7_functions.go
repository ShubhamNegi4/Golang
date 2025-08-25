package main

import( 
	"fmt"

)

//function printing Hello world
func sayhello(){
	fmt.Println("Hello world")
}
//function with return type int returing a+b
func add (a int, b int) int {
	return a+b
}

//func for swapping the values of x and y as a reference
func swap_with_reference(x *int, y *int) {
	*x += *y;
	*y = *x-*y;
	*x = *x-*y;
}


//Slice

//A slice is a small value (header) that contains a pointer to an underlying array, a length and a capacity. 
// When you pass a slice to a function you pass a copy of that header — 
// but both headers point to the same underlying array, so modifications to the elements (e.g. a[0] = ...) are visible to the caller.
func swapslice(a[]int){
	temp := a[0];
	a[0] = a[len(a)-1]
	a[len(a)-1]= temp
}


//Array

//passed by value 
func swaparrayval(a[4]int){
	temp := a[0];
	a[0] = a[len(a)-1]
	a[len(a)-1]= temp
}
//passed by reference

func swaparrayref(a *[4]int){
	temp:= a[0]
	a[0] = a[len(a)-1]
	a[len(a)-1] = temp
}


//func returning multiple values
func mutliplevalue()(int, string, float32){
	return 1, "hello word", 32.31
}



func main(){
	// normal function call
	sayhello()

	// function call with arguments
	fmt.Println(add(2,3))

	x := 4;
	y := 9;
	//Before swapping
	fmt.Println(x,y)
	//After swapping
	swap_with_reference(&x,&y) // sending x and y as reference as go always passes function arguments by value
	fmt.Println(x,y)


	//with slice
	a := []int{1,2,3,4,5}
	swapslice(a);//swapping first and last element
	fmt.Println(a);


	//with array
	b :=[4]int{1,2,3,4}
	swaparrayval(b)
	fmt.Println(b) //output- > [1,2,3,4]
	swaparrayref(&b)
	fmt.Println(b) //output -> [4,2,3,1]


	//multiple returned values
	fmt.Println(mutliplevalue()) // output-> 1 hello world 32.31



}