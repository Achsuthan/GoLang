package main

import "fmt"

func main() {
	var i = -1
	for i < 20 {
		i++
		if i == 5 {
			fmt.Println("Value reached")
			continue
		} else {
			fmt.Println("Hello value ", i)
		}

	}
}
