package main

import "fmt"

func checkError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("error-nil")
	}
}
func foo() error {

	return fmt.Errorf("Custom Error")
}

func bar() error {
	return nil
}
func main() {
	e := foo()
	checkError(e)

	b := bar()
	checkError(b)
}
