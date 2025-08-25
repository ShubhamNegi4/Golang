//main is a special package name in Go: when a package is main and it contains a main() function, the result is an executable program.
// If the package is not main, Go builds a library instead.

package main

import "fmt"

func main() {
	fmt.Println("Hello World") // for printing with \n
	fmt.Print("Hello World")   // for printing without \n
}
