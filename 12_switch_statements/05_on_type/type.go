package main

import "fmt"

// switch on types
// -- normally we switch on value
// -- go allows you to switch on type

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}


func main() {
  SwitchOnType(7)
  SwitchOnType("james")
  var t = Contact{"yea", "yodle"}
  SwitchOnType(t)
  SwitchOnType(t.greeting)
  SwitchOnType(t.name)
  }
}
