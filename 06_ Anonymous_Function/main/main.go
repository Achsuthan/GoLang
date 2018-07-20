package main

import "fmt"

func main() {
	Function_1(10)

	function_2 := func(x int) {
		fmt.Println("Hello this is a Anonymous function")

	}
	function_2(10)
}

func Function_1(x int) int {
	fmt.Println("Hello this is a function")
	return x
}
