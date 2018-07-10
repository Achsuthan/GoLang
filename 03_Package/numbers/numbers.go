package numbers

import (
	"Achsuthan/03_Package/add"
	"Achsuthan/03_Package/substract"
	"fmt"
)

var Number1 int
var Number2 int

//function must have starting with upercase letter
func SetValues(x int, y int) {
	Number1 = x
	Number2 = y
}

//function must have starting with upercase letter
func PrintValues() {
	//print the values from add and substract

	fmt.Println("Adding two numbers ", add.Addnumbers(Number1, Number2))
	fmt.Println("Substract two numbers ", substract.Substractnumbers(Number1, Number2))
}
