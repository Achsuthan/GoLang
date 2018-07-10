package substract

import "fmt"

//we can access only this function inside the package
func welcomeMessage() {
	fmt.Println("Number adding started")
}

//function must have starting with upercase letter
func Substractnumbers(x int, y int) int {
	welcomeMessage()
	return (x - y)
}
