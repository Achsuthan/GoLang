package main

import "fmt"

func main() {
	var i = 0
	for i = 0; i < 10; i++ {
		fmt.Println("value in loop ", i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("value in loop ", j)
	}
}
