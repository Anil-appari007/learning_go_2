package main

import "fmt"

func main() {
	aar1 := [5]int{5, 4, 3, 2, 1}
	fmt.Println(aar1)

	var arr2 [3]int
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	fmt.Println(arr2)

	arr3 := [5]int{}
	fmt.Println(arr3)

	stArr1 := [2]string{"a", "b"}
	fmt.Println(stArr1)
	stArr2 := [3]string{}
	fmt.Println(stArr2)

	fArr1 := [3]float64{4.3, 4.5}
	fmt.Println(fArr1)

	bArr1 := [2]bool{true}
	fmt.Println(bArr1)

}
