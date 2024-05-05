package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64
	ch := make(chan int)
	var wg sync.WaitGroup

	for t := 1; t <= 5; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range ch {
				fmt.Println("T", t, "received", task)
				// time.Sleep(2 * time.Second)
				counter.Add(1)
			}
		}()
	}

	for j := 1; j <= 50; j++ {
		ch <- j
	}
	close(ch)
	wg.Wait()
	fmt.Println("Count:", counter.Load())
}
