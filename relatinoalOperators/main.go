package main

import "fmt"

func main() {
	a := 1
	b := 2

	if a == b {
		fmt.Println("a==b")
	} else {
		fmt.Println("not a==b")
	}

	if a != b {
		fmt.Println("a!=b")
	}

	if a > b {
		fmt.Println("a>b")
	}
	if a < b {
		fmt.Println("a<b")
	}

	if a < b && b > a {
		fmt.Println("a<b && b>a")
	}

	if a > b || a > 0 {
		fmt.Println("Or condition")
	}

}
