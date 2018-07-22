package main

import "fmt"

const KM = 1000

func main() {
	var meters float64
	fmt.Print("Enter your meeter value : ")
	fmt.Scan(&meters) //when reading the user's imput we need to have a memory address
	fmt.Println("Converted value in KM ", KM*meters)
}
