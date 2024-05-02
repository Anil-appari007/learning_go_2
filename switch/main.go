package main

import "fmt"

func main() {
	a := 9

	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 4:
		fmt.Println(4)
	}

	switch {
	case a > 0 && a < 5:
		fmt.Println("a>0 && a < 5")
	default:
		fmt.Println("default-case")
	}
}
