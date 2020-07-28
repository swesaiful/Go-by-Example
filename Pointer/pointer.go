// Basic Pointer example in GO
// A pointer is a variable that represents the address (rather than the value) 
// of a data item within the computer’s memory location.
// The address operator (&) returns the address of a variable/function
// The indirection operator (*) returns data at an address (dereferencing)
 
package main
 
import "fmt"
 
func main(){
	
    i, j  := 16, 256
    
    fmt.Printf("i value: %v\n", i)     // print i value
	fmt.Printf("i memory address, &i = %v\n", &i)  // print i address for i value
    fmt.Printf("i type is = %T\n", i)	// type of i
    fmt.Printf("&i type is = %T\n", &i)	// type of &i
 
    p := &i // point to i address
    fmt.Printf("\nreturn &i in p: %v\n", p) // print i address
    fmt.Printf("see &p = %v\n", &p) // print p address
    fmt.Printf("p type is = %T\n", p)	// type of p
    fmt.Printf("see i value in *p: %v\n", *p)  // print i value
    fmt.Printf("see &i in *p: %v\n", &*p) // print i address
    fmt.Printf("*p type is = %T\n", *p)	// type of *p
    fmt.Printf("&*p type is = %T\n", &*p)	// type of &*p
 
    *p = 32 // set i value through the pointer
 
    fmt.Printf("\nSee new i value = %v\n", i)    // print i new value
    fmt.Printf("check &i for 2nd time = %v\n", &i) // print i address and see no change
    fmt.Printf("check &i in p for 2nd time = %v\n", p) // print i address and see no change
    fmt.Printf("check &p for 2nd time = %v\n", &p) // print p address and see no change
    fmt.Printf("see i new value  in *p = %v\n", *p)    // print i new value
    fmt.Printf("check &i in *p for 2nd time = %v\n", &*p) // print i address
 
    p = &j
 
    fmt.Printf("\nj value: %v\n", j) // print j value
    fmt.Printf("j memory address, &j = %v\n", &j)  // print j address
    fmt.Printf("check p value (address): %v\n", p)  // print new address
    fmt.Printf("check &p for 3rd time = %v\n", &p) // print p address and see no change
    fmt.Printf("see j value in *p = %v\n", *p) // print j value
    fmt.Printf("check &j in *p = %v\n", &*p)   // print j address
 
    *p = *p / 64    // divide j value through the pointer
 
    fmt.Printf("\nSee new j value = %v\n",j) // print j new value
 
    fmt.Printf("return &j in p: %v\n", p) // print j address
    fmt.Printf("check &p for 4th time = %v\n", &p) // print p address and see no change
    fmt.Printf("see j new value in *p: %v\n", *p)  // print j new value
    fmt.Printf("see &j in *p for 2nd time: %v\n", &*p) // print j address

}
