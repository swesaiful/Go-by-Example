package main

import (
"fmt"
)
var result int

func main() {

	// Arithmetic operators
	a := 9
	b := 3
	
	result = a + b
	fmt.Printf("\nThe addition result, a+b is: %v", result)	// 12

	result = a - b
	fmt.Printf("\nThe subtraction result, a-b is: %v", result)	// 6

	result = a * b
	fmt.Printf("\nThe multiplication result, a*b is: %v", result) // 27

	result = a / b
	fmt.Printf("\nThe division result, a/b is: %v", result)	// 3

	result = a % b
	fmt.Printf("\nThe modulus/mod result, a%%b is: %v", result)	// 0
	// fmt.Println("\nThe modulus result, a % b is: ", result)

	// Relational Operators
	c := 9
	d := 3
	
	fmt.Printf("\nRelational Operators, a>b is: %v", c>d)		// true
	fmt.Printf("\nRelational Operators, a<b is: %v", c<d)		// false
	fmt.Printf("\nRelational Operators, a>=b is: %v", c>=d)		// true
	fmt.Printf("\nRelational Operators, a>=b is: %v", c<=d)		// false

	// Equality Operators
	g := 9
	h := 3
	
	fmt.Printf("\nEquality Operators, g==h is: %v", g==h)		// false
	fmt.Printf("\nEquality Operators, g!=h is: %v", g!=h)		// true

	// Logical AND Operators
	i := true
	j := true
	k := true
	l := false
	
	fmt.Printf("\nLogical AND Operators, i&&j is: %v", i&&j)	// true
	fmt.Printf("\nLogical AND Operators, k&&l is: %v", k&&l)	// false

	// Logical OR Operators
	m := true
	n := true
	o := true
	p := false
	
	fmt.Printf("\nLogical OR Operators, m||n is: %v", m||n)	// true
	fmt.Printf("\nLogical OR Operators, o||p is: %v", o||p)	// true

	// Logical NOT Operators
	var x, y bool = true, true
	z := (x!=y)
	
	fmt.Printf("x != y is: %v", z)		// false

	// Bitwise AND Operators
	var q uint = 60				// 60 = 0011 1100
   	var r uint = 13				// 13 = 0000 1101
   	var s uint = 0          

  	s = q & r       			// 12 = 0000 1100
	fmt.Printf("\nBitwise AND Operators, the value of s is: %v", s)

	// Bitwise OR Operators
	var t uint = 60			// 60 = 0011 1100
   	var u uint = 13			// 13 = 0000 1101
   	var v uint = 0          

  	v = t | u       			// 61 = 0011 1101
	fmt.Printf("\nBitwise OR Operators, the value of v is: %v", v)

	// Bitwise XOR Operators
	var aa uint = 60			// 60 = 0011 1100
   	var bb uint = 13			// 13 = 0000 1101
   	var cc uint = 0          

  	cc = aa ^ bb       			// 49 = 0011 0001
	fmt.Printf("\nBitwise XOR Operators, the value of cc is: %v", cc)

	// Shift Left Operators
	var dd uint = 60			// 60 = 0011 1100
   	var ee uint = 0          

  	ee = dd << 2      			// 240 = 1111 0000
	fmt.Printf("\nShift Left Operators, the value of ee is: %v", ee)

	// Shift Right Operators
	var gg uint = 60			// 60 = 0011 1100
   	var hh uint = 0          

  	hh = gg >> 2      			// 15 = 0000 1111
	fmt.Printf("\nShift Right Operators, the value of hh is: %v", hh)

}
