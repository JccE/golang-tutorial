package main

import "fmt"

func main() {
	if true {
		fmt.Println("this ran because of true")
	}
	if false {
		fmt.Println("this did not run")
	}
}
