package main

import (
	"fmt"
)

func main() {
	str1 := "sampleData"
	fmt.Println(str1)

	for _, each := range str1 {
		fmt.Println(each)
		fmt.Println(string(each))
	}
}
