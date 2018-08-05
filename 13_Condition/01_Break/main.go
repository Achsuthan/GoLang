package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println("hello value ", i)
		if i == 5 {
			fmt.Println("value reached")
			break
		}
	}
}
