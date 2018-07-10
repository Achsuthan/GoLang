package main

import (
	"Achsuthan/03_Package/numbers"
	"Achsuthan/03_Package/stringutil"
	"fmt"
)

func main() {
	fmt.Println("hello Package")
	//access the variable form the package
	fmt.Println("hello ", stringutil.UserName)

	//access the function form the package
	numbers.SetValues(20, 30)
	numbers.PrintValues()

}
