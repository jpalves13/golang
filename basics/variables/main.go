package main

import "fmt"

// Remember constant is immutable
const (
	// public const: uppercase in first name
	PublicConst = "a"
	// private const: lowercase in first name
	privateConst = "b"
)

var (
	// public variable: uppercase in first name
	PublicVar = "a"
	// private variable: lowercase in first name
	privateVar = "b"
)

func main() {
	//bad way
	var a string = "a"
	fmt.Print(a)

	//good way
	b := "b"
	fmt.Print(b)
}
