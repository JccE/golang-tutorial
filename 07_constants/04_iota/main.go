package main

import "fmt"

const (
	A = iota
	B
	C
)

const (
	D = iota
	E
	F
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
