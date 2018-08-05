package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Println("Hello Value ", j)
		}
		fmt.Println()
	}
}
