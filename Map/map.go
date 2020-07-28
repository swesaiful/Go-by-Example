package main

import "fmt"

func main() {
	var m map[string]int
	// create a map using a make keyword
	m = make(map[string]int)
	fmt.Printf("\nInitial map: %v\n", m)
	// insert value into the map
	m["key"] = 42
	fmt.Printf("Map with value: %v\n", m)
	delete(m, "key")
	fmt.Printf("After deleting the value: %v\n", m)
	
	// declare and initialize map
	var idMap = map[string]int {
		"a": 123,
		"b": 371,
		"c": 341,
	}
	fmt.Printf("This is another map: %v\n", idMap)
	// len function will help us to know about the number of 
	// key-value pairs inside the map
	fmt.Printf("The length of the map: %v\n", len(idMap))
	// access the elements in a map
	fmt.Printf("The b is: %v\n", idMap["b"])

	// adding a key/value pair to the map
	idMap["d"] = 456
	fmt.Printf("After adding value to the map: %v\n", idMap)

	// two-value assignment tests for existence of the key
	// id is the value, ok is presence of key and its Boolean
	id, ok := idMap["b"]
	// ok is going to be true if the key "b" is present in the map
	fmt.Printf("Is the key %v is in the map? Answer = %v\n", id, ok)

	// Use a for loop with the range keyword
	for key, val := range idMap{
		fmt.Printf("Key = %v, Value = %v\n", key, val)
	}
}
