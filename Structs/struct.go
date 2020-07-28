package main

import (
	"fmt"
)

// Person is a name
type Person struct {
	name    string
	address string
	Phone   int
}

func main() {

	//creating struct specifying field names
	
	p1 := Person{
		name:    "Adam",
		address: "Skyview 22",
		Phone:   4601234567,
	}

	//creating struct without specifying field names
	p2 := Person{"Sana", "Parklane 33", 46234567}

	fmt.Printf("Person 1: %v\n", p1)
	fmt.Printf("Person 2: %v\n", p2)
	fmt.Printf("\nPerson1.address: %v\n", p1.address)
	fmt.Printf("Person2.phone: %v\n", p2.Phone)

}
