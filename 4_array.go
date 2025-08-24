package main

import(
	"fmt"
	"sort"

)

func main(){

	//static array
	var arr[4] int = [4] int {1,2,3,4}
	var arr2 = [4] int {6,7,8}
	fmt.Println(arr)
	fmt.Println(arr[0])
	arr[0] = 5
	fmt.Println(arr[0])
	fmt.Println(arr2)

	//slices/dynamic array
	var temp = []int{1,2,3,4,5}
	fmt.Println(temp)
	temp = append(temp, 6)
	fmt.Println(temp);
	fmt.Println(len(temp))


	//slice ranges
	brr1 := arr[1:3]
	brr2 := arr[:4]
	brr3 := arr[0:]
	fmt.Println(brr1,brr2,brr3)



	//sort an array using sort func
	var crr1 = []int{3,1,3,5,6,1,6,8,3,4}
	sort.Ints(crr1)
	fmt.Println(crr1)


}