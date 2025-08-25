package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Scan(&str)
	err := os.WriteFile("file.txt", []byte(str), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File saved")

	data,err := os.ReadFile("file.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

}

/*

output

shubhamnegi@senkus-MacBook-Air Golang % go run 12_readwritefile.go
Hello World    
File saved
[72 101 108 108 111 32 87 111 114 108 100 10]
Hello World

shubhamnegi@senkus-MacBook-Air Golang % 


*/