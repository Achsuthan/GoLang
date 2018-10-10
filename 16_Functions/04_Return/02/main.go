package main

import "fmt"

func main() {
	fmt.Println("Hello Main")
	greet("Mahendran", "Achsuthan")
}

//this is called return naming
func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
