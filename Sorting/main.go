package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"z", "y", "x", "a"}
	fmt.Println(letters)

	slices.Sort(letters)
	fmt.Println(letters)

	nums := []int{9, 5, 8, 6}
	fmt.Println("nums before sort", nums)
	slices.Sort(nums)
	fmt.Println(nums)
	fmt.Println("Numbers are sorted:", slices.IsSorted(nums))

}
