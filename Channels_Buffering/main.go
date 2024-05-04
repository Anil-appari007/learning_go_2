package main

import "fmt"

func main() {
	ch1 := make(chan string)
	info := ""

	go func() { ch1 <- "data1" }()
	info = <-ch1
	fmt.Println("info", info)
	go func() { ch1 <- "data2" }()
	info = <-ch1
	fmt.Println("info", info)

	ch2 := make(chan string, 2)
	// info2 := ""
	go func() { ch2 <- "info1" }()
	// info = <-ch2
	// fmt.Println("CH2 -", info)
	fmt.Println(<-ch2)

	go func() { ch2 <- "info2" }()
	// info = <-ch2
	// fmt.Println("CH2 -", info2)
	fmt.Println(<-ch2)

	go func() { ch2 <- "info3" }()
	fmt.Println(<-ch2)
}
