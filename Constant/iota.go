package main

import(
	"fmt"
)

// iota represents successive untyped integer constants 0, 1, 2
// Its value starting at zero. 
// It can be used to construct a set of related constants

const (
	x = iota
	y
	z
)

func main(){
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}