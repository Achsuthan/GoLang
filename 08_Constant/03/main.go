package main

//Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers. Because it can be used in expressions, it provides a generality beyond that of simple enumerations.
//it will reset call it in another const
import "fmt"

const (
	a = 10
	b
	c
)
const (
	d = iota
	e
	f
	g
)
const (
	h = iota
	i
)

func main() {
	fmt.Println("Hello value ", d)
	fmt.Println("Hello value ", e)
	fmt.Println("Hello value ", h) //for h it is reset to zero
	fmt.Println("Hello value ", g) //g is declared berfor the h constant
}
