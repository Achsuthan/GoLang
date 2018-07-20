package main

import "fmt"

func main() {
	wra := wrapper()
	fmt.Println(wra())
	fmt.Println(wra())

}

//this warpper function will return a function which will ruturn a int value
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
