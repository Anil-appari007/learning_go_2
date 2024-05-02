package main

import "fmt"

func main() {

	type User struct {
		Name string
		Age  int
	}

	a := User{Name: "a", Age: 1}
	fmt.Println(a)

	var b User
	b.Name = "b"
	b.Age = 2
	fmt.Println(b)

	// anonymous struct
	c := struct {
		User string
	}{
		User: "c",
	}
	fmt.Println(c)

}
