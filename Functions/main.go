package main

import "fmt"

func sayHello(s string) {
	fmt.Println(s)
}

func addNum(a, b int) (int, error) {
	c := a + b
	return c, nil
}

func addNumInfy(nums ...int) int {
	c := 0
	for _, num := range nums {
		fmt.Println(num)
		c += num
	}
	return c
}

func recFunc(a int) {
	if a > 0 && a != 0 {
		fmt.Println("n:", a)
		a--
		recFunc(a)
	} else if a == 0 {
		fmt.Println("n:", a)
	}
}

func main() {
	sayHello("a")
	sayHello("b")

	fmt.Println(addNum(1, 2))
	fmt.Println(addNumInfy(1, 2, 3, 4))

	recFunc(4)

}
