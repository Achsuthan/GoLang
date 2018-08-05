package main

import "fmt"

func main() {
	switch "Achsuthan" {
	case "Tim":
		fmt.Println("Hello Tim")
	case "Johon":
		fmt.Println("Hello Johon")
	case "Maria":
		fmt.Println("Hello Maria")
	default:
		fmt.Println("User not found")
	}

}
