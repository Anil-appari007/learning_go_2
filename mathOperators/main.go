package main

import "fmt"

func main() {
	fmt.Println("math operators")
	a := 1
	b := 2

	c := a + b
	d := a - b
	e := a / b
	f := a * b

	fmt.Print(a, b, c, d, e, f)

	fmt.Println("\nModulus")
	h := b % a
	fmt.Println(h)
}
