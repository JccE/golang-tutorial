package main

import (
	"fmt"
)

func main() {
	switch "pems" {
	case "articles", "notes":
		fmt.Println("a in ns")
	case "boo", "sns":
		fmt.Println("boo in sns")
	case "pics", "pems":
		fmt.Println("psics in pens")
	}
}
