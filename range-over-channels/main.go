package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	// goroutine 1 [chan receive]:
	// If channel is not closed
	close(ch)
	fmt.Println("Length of ch", len(ch))

	for each := range ch {
		fmt.Println(each)
	}
	fmt.Println("Length of ch", len(ch))
}
