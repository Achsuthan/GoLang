package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var user = "tim"
	switch user {
	case "Tim", "Jamal":
		fmt.Println("Hello Tim or Jamal")
	case "Johon":
		fmt.Println("Hello Johon")
	case "achsuthan":
		fmt.Print("hello achsuthan")
	default:
		fmt.Println("User not found")
	}

}
