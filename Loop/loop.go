// There is only for loop in Go, no While or DO loop

package main
import "fmt"
	func main() {
	
	i := 1
	for i <= 3 {
	fmt.Printf("i = %v\n",i)
	i = i + 1
	}
	
	fmt.Print("\n")

	for j := 1; j <= 3; j++ {
		fmt.Printf("j = %v\n", j)
	}


}