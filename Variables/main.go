package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a)
	fmt.Println(b)

	var c int = 3
	fmt.Println(c)

	d, e := 4, 5
	fmt.Println(d)
	fmt.Println(e)

	var (
		f int    = 6
		g string = "g"
	)

	fmt.Println(f)
	fmt.Println(g)

	var h, i int = 9, 10
	fmt.Println(h)
	fmt.Println(i)

	var j int
	fmt.Println(j)

	var k string
	fmt.Println(k)
}
