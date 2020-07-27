package main

import "fmt"

// this is a comment

func main() {
	switch n := 6; n {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	case 4:
		fmt.Println("FOUR")
	case 5:
		fmt.Println("FIVE")	
	default: // default case
		fmt.Println("Incorrect number.")

	}
}
