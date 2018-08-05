package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var user = "Achsu"
	switch {
	case user == "Tim", user == "Jamal":
		fmt.Println("Hello Tim or Jamal")
	case user == "Johon":
		fmt.Println("Hello Johon")
	case user == "Achsu", len(user) > 2:
		fmt.Println("hello achsuthan")
	default:
		fmt.Println("User not found")
	}

}
