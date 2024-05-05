package main

import (
	"fmt"
	"sync"
	"time"
)

func sleeper(wg *sync.WaitGroup, n int) {
	fmt.Println("Sleeping for", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Sleep is done for", n)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for each := range 5 {
		wg.Add(1)
		go sleeper(&wg, each)
	}

	fmt.Println("Waiting for routines to complete")
	wg.Wait()
	fmt.Println("Routines are completed")
}
