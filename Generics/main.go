package main

import "fmt"

func addInt(a, b int) (int, error) {
	return a + b, nil
}

func addFloats(a, b float64) (float64, error) {
	return a + b, nil
}

func addNums[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Generics")
	sum1, err := addInt(1, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("sum1 = %d", sum1)
	}

	float1, err := addFloats(1.5, 2.5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\nFloat1 = %v", float1)
	}

	fmt.Println("\nnums1 =", addNums(1, 2))
	fmt.Println("\nnums1 =", addNums(1.22, 2.43))

}
