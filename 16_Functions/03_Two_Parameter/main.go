package main

import "fmt"

func main() {
	fmt.Println("hello Main")
	multipleParameter("Mahendran ", "Achsuthan")
}

//we can use comma between identifers and put the identifier type
func multipleParameter(firstname, lastname string) {
	//we can print two values with comma
	fmt.Println(firstname, lastname)
}
