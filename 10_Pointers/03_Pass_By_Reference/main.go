package main

import "fmt"

func foo(x *int) {
	fmt.Println("Memory address of x in function ", &x)
	fmt.Println("Vlaue of x in function ", x)
	*x = 100
}
func main() {
	x := 10
	fmt.Println("Memory address of x in main ", &x)
	fmt.Println("Value of x in main ", x)

	foo(&x)

	fmt.Println("Memory address of x in main after call the foo function ", &x)
	fmt.Println("Vlaue of x after call call the foo function in main ", x)

}
