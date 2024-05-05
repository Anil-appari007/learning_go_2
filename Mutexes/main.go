package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for each := range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(each)
			time.Sleep(3 * time.Second)
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("count", counter)
}
