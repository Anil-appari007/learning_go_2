package main

import (
	"fmt"
	"reflect"
)

type geometry interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	r float32
}

func (c Circle) Area() float32 {
	return 3.14 * c.r * c.r
}

func (c Circle) Perimeter() float32 {
	return 2 * 3.14 * c.r
}

type Rectangle struct {
	l float32
	b float32
}

func (r Rectangle) Area() float32 {
	return r.l * r.b
}

func (r Rectangle) Perimeter() float32 {
	return 2 * r.l * r.b
}

func main() {
	o1 := Circle{r: 4}
	fmt.Printf("Type: %T", o1)
	fmt.Println(reflect.TypeOf(o1))
	fmt.Println("Area", o1.Area())
	fmt.Println("perimeter", o1.Perimeter())

	o2 := Rectangle{l: 3.5, b: 5}
	fmt.Println(o2.Area())
	fmt.Println(o2.Perimeter())

	type square struct {
		s float32
	}
	o3 := square{s: 4.5}
	fmt.Println(reflect.TypeOf(o3))
}
