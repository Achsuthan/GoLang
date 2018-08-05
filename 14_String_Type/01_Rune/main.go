package main

import "fmt"

func main() {
	fmt.Println("hello")

	for i := 0; i < 150; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	var value = 'a' //this is a rune in Golang

	fmt.Printf("%T ", value)
	fmt.Println()
	fmt.Printf("%v", value)
	fmt.Println()
	fmt.Printf("%T ", rune(value))
	fmt.Println()
}
