package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (person *Person) GetFullName() string {
	return fmt.Sprintf(person.FirstName + person.LastName)
}

func main() {
	var a Person
	a.FirstName = "b"
	a.LastName = "c"
	fmt.Println(a.GetFullName())

}
