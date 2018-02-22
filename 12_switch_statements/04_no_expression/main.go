package main

import (
	"fmt"
)

func main() {
	theName := "pems"

	switch {
	case len(theName) == 2:
		fmt.Println("a in ns")
	case theName == "boo":
		fmt.Println("boo in sns")
	case theName == "pics":
		fmt.Println("psics in pens")
	default:
		fmt.Println("nothing matches")
	}
}
