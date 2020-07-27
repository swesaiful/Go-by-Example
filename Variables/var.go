package main

import (
	"fmt"
)

// global variable declaration
var g = "Go"

func main(){
	// local variable declaration
	var m, a, rn string = "Gopher", "November 10, 2009","Golang"
	u := 9.4			
	l := true				
	const Truth = true		// constant variable declaration

	fmt.Printf("\n%s is designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.", g)
	fmt.Printf("\n%s is an iconic mascot of Go.", m)
	fmt.Printf("\nGo first appears in %s.", a)
	fmt.Printf("\nIt is often referred to as %s due to its domain name, but the proper name is Go.", rn)
	// string concatenation
	fmt.Printf("\n\nAccording to the Stack Overflow Developer Survey 2020, " + 
					"\n%g percent professional developers currently using go. ", u)
	fmt.Println("\nPeople loves Go?", l)
	// string concatenation
	fmt.Println("\nGo is an imperative, statically typed, compiled,\ngeneral-purpose, open source, " +
	 "concurrent " + "\nand a modern " + " programming language.")
	
	fmt.Println("\nGo rules?", Truth)
}