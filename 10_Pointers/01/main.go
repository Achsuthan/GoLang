package main

import "fmt"

func main() {
	x := 12
	var b *int = &x
	fmt.Println("Value of x is ", x)
	fmt.Println("Memory address of x ", &x)
	fmt.Println("Memory address of b ", &b)
	fmt.Println("Vlaue in b ", b)
	fmt.Println("Pointing value of b is ", *b)

}
