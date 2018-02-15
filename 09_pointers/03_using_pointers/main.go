package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20919a22

	var b *int = &a
	fmt.Println(b)  // 0x20919a22
	fmt.Println(*b) // 43

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) //42

	// this is useful
	// we can pass a memory addrewss instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whetever is stored at that memory address
	// this makes our programs more performant
	// we dont have to pass around large amounts of data

	// everything is PASS BY VALUE
	// when we pass a memory address, we are passing a value
}
