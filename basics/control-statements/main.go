package main

import "fmt"

func main() {
	boolean()
	operators()
	ifElse()
	condSwitch()
}

func boolean() {
	a := 10
	b := 8

	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a == b) // false
}

func operators() {
	a := 8
	b := 10

	fmt.Println(a == b) // false
	fmt.Println(a == a) // true
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // true
	fmt.Println(a <= b) // true
	fmt.Println(a <= a) // true
	fmt.Println(a >= a) // true
}

func ifElse() {
	// simple
	if 10 > 8 {
		fmt.Println("yes") // printed
	} else {
		fmt.Println("no")
	}

	// more conditionals
	if 10 > 8 {
		fmt.Println("yes") // printed
	} else if 10 == 8 {
		fmt.Println("no")
	} else {
		fmt.Println("no")
	}

	// init
	if result := zero(); result == 0 {
		fmt.Println("yes") // printed
	}

}

func zero() int {
	return 0
}

func condSwitch() {
	switch 1 {
	case 1:
		fmt.Println("Yes is 1") // printed
	case 2:
		fmt.Println("Yes is 2")
	case 3:
		fmt.Println("Yes 3")
	default:
		fmt.Println("Nothing")
	}
}
