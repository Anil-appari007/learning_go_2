package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// var count uint64
	var count atomic.Int64
	var wg sync.WaitGroup

	for i := 1; i <= 25; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println("Thread", i)
			time.Sleep(2 * time.Second)
			// atomic.AddUint64(&count, 1)
			count.Add(1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// fmt.Println("Count", count)
	fmt.Println("Count", count.Load())
}
