package main

import "fmt"

func main() {

	fmt.Println("Hello main")
	fmt.Println(greet("Mahendran ", "Achsuthan"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
