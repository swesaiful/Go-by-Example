package main

import (
	"fmt"
)

func main() {  
	//Boolean data types
    x := true
    y := false
    fmt.Println("x:", x, "y:", y)
    i := x && y
    fmt.Println("i:", i)
    j := x || y
	fmt.Println("j:", j)
	
	//Integer data types
	a := 5
	fmt.Println("value of a is:", a)

	//Floating-point data types
	b := 10.5
	fmt.Println("value of b is:", b)

	//Complex data types
	c1 := complex(1, 3)
    c2 := 4 + 7i
    cAdd := c1 + c2
    fmt.Println("sum of c1 and c2 is:", cAdd)
    cMul := c1 * c2
	fmt.Println("product of c1 and c2 is:", cMul)
	
	//String data types
	n := "Go"
    m := "Gopher"
    name := n +" "+ m
	fmt.Println("Go's mascot name is ",name)

	//Type Conversion
	g := 5      		// g - int data type
    h := 15.50    		// h - float64 data type
    sum := g + int(h) 	// h is converted to int
    fmt.Println("This result is after the type conversion: ", sum)
}