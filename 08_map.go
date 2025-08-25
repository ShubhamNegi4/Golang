package main
import (
	"fmt"
)

func main(){
//	mp := map[key]value{key1 : value1}
	mp:= map[int]int{
		1 :3,
		2 :4,
		3 :5,
	}
	fmt.Println(mp) //output -> map[1:3 2:4 3:5]
	fmt.Println(mp[1]) //output -> 3
	mp[1] = 8

	fmt.Println(mp) //output -> map[1:8 2:4 3:5]
}