package main

import "fmt"

func main() {
	switch "Jamal" {
	case "Tim":
		fmt.Println("hello Time")
		fallthrough
	case "Johon":
		fmt.Println("hello Johon")
	case "Jamal":
		fmt.Println("hello jamal")
	default:
		fmt.Println("User not found")
	}

}
