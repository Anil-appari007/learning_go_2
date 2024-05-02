package main

import "fmt"

func addNum[T int | float32 | float64](a []T) T {
	var Sum T
	for _, each := range a {
		Sum += each
		// fmt.Println(each)
	}

	// fmt.Println(Sum)
	return Sum
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(addNum(a))

	b := []float32{1, 2.3, 3.4}
	fmt.Println(addNum(b))
}
