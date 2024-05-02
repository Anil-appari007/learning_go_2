package main

import "fmt"

func WithoutPointer(n int) {
	n = 0
}

func WithPointer(n *int) {
	fmt.Println("Withpointer-received value", *n)
	*n = 0
	fmt.Println("After-update with-in-func", *n)
}

func main() {
	a := 1
	fmt.Println("Initial_Value of a:", a)

	WithoutPointer(a)
	fmt.Println("after WithoutPointer", a)

	WithPointer(&a)
	fmt.Println("After WithoutPointer in main a:", a)
}
