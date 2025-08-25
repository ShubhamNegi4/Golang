// Go does not have classes like Java, C++, or Python.

// But that doesn’t mean Go cannot do object-oriented programming (OOP). Instead of classes, Go provides a different model using:
package main

import (
	"math"
)

type bill struct {
	name string
	tip float32
	number int
}


func newbill(name string) bill {
	b:=  bill{
		name : name,
		tip : math.Pi,
		number : 000000,
	}
	return b
}