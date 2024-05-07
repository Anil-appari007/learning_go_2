package main

import (
	"fmt"
	"time"
)

func exitMsg() {
	fmt.Println("Program is exiting.")
}

func sleep() {
	num := 10
	fmt.Printf("Sleeping for %vs\n", num)
	time.Sleep(time.Duration(num) * time.Second)
}

func main() {
	defer exitMsg()
	fmt.Println("Sample Defer")
	sleep()
}
