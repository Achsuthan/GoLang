package main

import "fmt"

var x int = 12

//function or variable name start with lowercase can accessible inside the package only start with upercase is accessible outside the package

func main() {
	fmt.Println("Value of X ", x)
	foo()
}

func foo() {
	fmt.Println("Value fo X in function ", x)
}
