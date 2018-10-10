package main

import "fmt"

func main() {
	fmt.Println("Main function")
	testFunction("Achsuthan")
}

// name is identifier string is the type of the identifier
func testFunction(name string) {
	fmt.Println("Hello " + name)
}
