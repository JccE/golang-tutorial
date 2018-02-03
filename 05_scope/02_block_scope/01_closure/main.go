// closure: x is declared outside main and still works be cause it has package level scope
package main

import (
	"fmt"
)

func main() {
	fmt.Println(x)
}

var x int = 42
