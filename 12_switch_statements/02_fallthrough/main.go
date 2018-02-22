package main

import "fmt"

func main() {
	switch "will" {
	case "bill":
		fmt.Println("billy")
	case "jim":
		fmt.Println("jimmy")
	case "will":
		fmt.Println("willy")
		fallthrough
	case "ho":
		fmt.Println("hoes")
		fallthrough
	case "loft":
		fmt.Println("lofts")
	case "sally":
		fmt.Println("sally")
	}
}
