package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func(ch chan int) {
		time.Sleep(2 * time.Second)
		ch1 <- 111

	}(ch1)

	select {
	case data := <-ch1:
		fmt.Println("data:", data)
	case <-time.After(3 * time.Second):
		fmt.Println("TIme-Out")
	}

}
