package main

import "fmt"

func main() {
	a := 10
	b := "hello"
	c := false
	d := 10.0

	//%T will give the type of the variable
	fmt.Printf("Values %T \n", a)
	fmt.Printf("Values %T \n", b)
	fmt.Printf("Values %T \n", c)
	fmt.Printf("Values %T \n", d)
}
