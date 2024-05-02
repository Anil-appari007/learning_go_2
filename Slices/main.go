package main

import "fmt"

func main() {
	var s1 []int

	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
		fmt.Println(i)
	}
	fmt.Println(s1)

	for j := range len(s1) {
		s1[j] += 5
	}
	fmt.Println(s1)

	s2 := make([]int, 0, 5)
	for k := range 27 {
		s2 = append(s2, k)
		fmt.Println("Len:", len(s2))
		fmt.Println("Cap", cap(s2))
	}

	fmt.Println(s2)

	fmt.Println(s2[1:6])
	fmt.Println(s2[3:])
	s3 := make([]int, 0, 3)
	fmt.Println("s3", s3)
	fmt.Println(len(s3))

	s4 := []int{}
	fmt.Println("s4", s4)
}
