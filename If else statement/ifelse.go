package main
import "fmt"

// this is a comment

var studentGrade int = 80

func main() {	
	// if statement
	if studentGrade >= 60 {
	fmt.Println(`Congratulation! You are "Passed".`)
	}

	// if statement with initialization
	if x := 100; x == 100 {
		fmt.Println("Gopher")
	}
	
	// if-else Statement
	if studentGrade >= 60 {
		fmt.Println(`Congratulation! You are "Passed".`)
	} else {
		fmt.Println(`Sorry! You are "Failed".`)
	}

	// if-else if-else Statement
	if studentGrade >= 90 {
		fmt.Println( `Your result is "A"` )
	} else if studentGrade >= 80 {
		fmt.Println( `Your result is "B"` )
	} else if studentGrade >= 70 {
		fmt.Println( `Your result is "C"` )
	} else if studentGrade >= 60 {
		fmt.Println( `Your result is "D"` )
	} else {
		fmt.Println( `Sorry! You are "Failed"` )
	}

	// If statement with return

	num := 10;
    if num%2 == 0 { 	//checks if number is even
        fmt.Println(num, "is even")
        return
    }
    fmt.Println(num, "is odd")
}
