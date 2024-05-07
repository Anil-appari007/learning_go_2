package main

import "fmt"

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Error: ", r)
			fmt.Println("Recovering")
		}
	}()
	panic("Manual Panic")

}
