package main

import (
	"fmt"
)

func main() {
	// array declaration
	// arr := [...] string {"a", "b", "c", "d", "e", "f", "g"}

	// slice declaration
	sli := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Printf("The new slice sli is: %v\n", sli)
	s1 := sli[1:3]
	s2 := sli[2:5]
	fmt.Printf("\nThe value of s1: %v", s1)
	fmt.Printf("\nThe value of s2: %v\n", s2)
	fmt.Printf("S1: len = %v, cap = %v\n", len(s1), cap(s1))
	fmt.Printf("S2: len = %v, cap = %v\n", len(s2), cap(s2))

	// create a new slice by using append function
	s3 := append(sli, "h", "i")
	fmt.Printf("\nUsing append, the slice s3 is: %v\n", s3)

	// create a copy of a slice by using copy function
	s4 := []int{1, 2, 3}
	s5 := make([]int, 2) 		// make function
	//s5 := new([2] int) [0:2]	// new function
	copy(s5, s4)
	fmt.Printf("\nThe s4 slice :%v, and the s5 slice :%v\n", s4, s5)

}
