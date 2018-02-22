package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("dan")
	case "Medhi":
		fmt.Println("med")
	case "Jenny":
		fmt.Println("jen")
	default:
		fmt.Println("Bitches")
	}
}
