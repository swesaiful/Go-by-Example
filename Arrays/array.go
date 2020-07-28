package main
import "fmt"

func main() {
	// Declare and initialize array
	// var x [11]float64 = [...] float64 {0, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// Using shorthand array declaration
	// x := [...] float64 {0, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// This is an alternative way to declare an array
	var x [11]float64
	fmt.Printf("x: %v\n", x)
	// Insert value into the array x
	x[0] = 0
	x[1] = 2
	x[2] = 4
	x[3] = 8
	x[4] = 16
	x[5] = 32
	x[6] = 64
	x[7] = 128
	x[8] = 256
	x[9] = 512
	x[10] = 1024
	
	var total float64
	/*
	for i := 0; i < len(x); i++ {
	total += x[i]
	fmt.Printf("total= %v", total)
	} 
	*/
	for _, value := range x{
		// By using blank identifier (_) to tell the compiler that 
		// we don't need the iterator variable.

		total += value
		fmt.Printf("\ntotal= %v", total)
	}
	fmt.Printf("\nThe total is divided by the length of the array: %v", total / float64(len(x)))
		/*
			 Here, total is a float64 while len(x) is an int data type. 
			 So we need to convert len(x) into a float64.
		*/
}
