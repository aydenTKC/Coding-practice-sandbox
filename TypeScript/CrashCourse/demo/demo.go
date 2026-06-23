package main

import (
	"fmt"
)

func main() {
	// Declare a variable with a specific type
	var myVariable int = 10
	fmt.Println(myVariable)

	// Attempting to change the type myVariable will resuilt in a comile-time error
	// myVariable = "Hello, World!" // This line would cause an error if uncommented

	// Correct way to declare a new variable with a different type
	var myOtherVariable string = "Hello World!"
	fmt.Println(myOtherVariable)

}
