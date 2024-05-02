package main

import "fmt"

type account struct {
	id int
}

type user struct {
	name string
	id   account
}

func (u user) Info() {
	fmt.Println("User: ", u.name)
	fmt.Println("Id: ", u.id)
}
func main() {
	a := user{
		name: "a",
		id:   account{id: 1},
	}

	fmt.Println(a)
	a.Info()
}
