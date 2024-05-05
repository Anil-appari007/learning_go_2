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

	for t := 1; t <= 15; t++ {
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

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex
// 	count := 0
// 	for i := 1; i <= 50; i++ {
// 		wg.Add(1)
// 		go func(wg *sync.WaitGroup, mutex *sync.Mutex, count *int, i int) {
// 			defer wg.Done()
// 			fmt.Println(i)
// 			// time.Sleep(1 * time.Second)
// 			mutex.Lock()
// 			*count++
// 			mutex.Unlock()
// 		}(&wg, &mutex, &count, i)
// 	}
// 	wg.Wait()
// 	total := &count
// 	fmt.Printf("Count: %d", *total)
// }
