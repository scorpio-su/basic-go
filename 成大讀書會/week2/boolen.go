package main

import "fmt"

func main() {

	var b bool //init false

	//b = 1 // error
	//b = bool(0) // error

	b = (1 != 0) // correct
	fmt.Println(b)
}
