package main

import "fmt"

func main() {
	s := "CCNSccns"
	fmt.Println(len(s))     // "8"
	fmt.Println(s[0], s[4]) // "67 99" ('C' and 'c')

	c := s[:len(s)] // panic: index out of range
	fmt.Println(c)

	fmt.Println("Hello, " + s[:4]) // "Hello, CCNS"

	//s[0] := 'A' // compile error: cannot assign to s[0]
}
