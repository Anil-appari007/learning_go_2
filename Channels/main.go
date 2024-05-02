package main

import (
	"fmt"
	"time"
)

func waitForAll(allData []int, ch chan int) {
	total := 0
	for _, each := range allData {
		fmt.Println("Sleeping for", each)
		time.Sleep(time.Duration(each) * time.Second)
		total++
	}
	ch <- total
}

func main() {
	fmt.Println("Channels")
	ch := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- 100
	}()
	v := <-ch
	fmt.Println("v =", v)

	waitTime := []int{1, 2, 3, 4}
	ch2 := make(chan int)
	go waitForAll(waitTime, ch2)

	totalValue := <-ch2
	fmt.Println("Total slept time", totalValue)
}
