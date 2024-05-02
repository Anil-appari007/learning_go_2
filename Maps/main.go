package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2

	fmt.Println(m1)
	fmt.Println(len(m1))

	delete(m1, "a")
	fmt.Println(m1)

	v, ok := m1["a"]
	fmt.Println(v, ok)

	v2, ok2 := m1["b"]
	fmt.Println(v2, ok2)

	for range m1 {
		fmt.Println(m1)
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	clear(m1)
	fmt.Println(m1)

}
