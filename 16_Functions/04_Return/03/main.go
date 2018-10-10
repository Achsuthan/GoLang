package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println(greet(" Mahendran", " Achsuthan"))
}

//multiple reuturn value function
func greet(fname, laname string) (string, string) {
	return fmt.Sprint(fname, laname), fmt.Sprint(laname, fname)
}
