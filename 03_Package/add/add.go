package add

import "fmt"

//we can access only this function inside the package
func welcomeMessage() {
	fmt.Println("Number adding started")
}

//function must have starting with upercase letter; if its in lower case we can access inside the package only
func Addnumbers(x int, y int) int {
	welcomeMessage()
	return (x + y)
}
