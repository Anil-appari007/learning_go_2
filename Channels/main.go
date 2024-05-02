package main

import (
	"fmt"
	"sync"
	"time"
)

/*

 */

func testFunc(ch chan int, count int) {
	count++
	ch <- count
}

// func sleep(item string, indexOfItem int, tL int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(item, indexOfItem)
// 	time.Sleep(4 * time.Second)
// 	fmt.Printf("sleep done for %s. Count:%d/%d \n", item, indexOfItem, tL)
// }

func sleep(item string, indexOfItem int, tL int, wg *sync.WaitGroup, count *int, m *sync.Mutex) {
	defer wg.Done()
	fmt.Println(item, indexOfItem)
	time.Sleep(4 * time.Second)
	m.Lock()
	*count++
	m.Unlock()
	// fmt.Printf("sleep done for %s. Count:%d/%d \n", item, indexOfItem, tL)
	fmt.Printf("sleep done for %s. Count:%d/%d \n", item, *count, tL)

}

// func sleep(item string, indexOfItem int, tL int, wg *sync.WaitGroup, ch chan int) {
// 	defer wg.Done()
// 	fmt.Println(item, indexOfItem)
// 	time.Sleep(4 * time.Second)
// 	tL--
// 	ch <- tL
// 	fmt.Printf("sleep done for %s. Count:%d/%d \n", item, indexOfItem, tL)
// }

func main() {
	ch := make(chan int)
	var count int
	// ch := new(chan int)
	go testFunc(ch, count)
	a := <-ch

	go testFunc(ch, a)
	a = <-ch

	fmt.Println(a)

	// iCount := 0
	// cCount := 0

	items := [4]string{"a", "b", "c", "d"}
	var wg sync.WaitGroup
	// ch2 := make(chan int)
	tL := len(items)
	// for n, each := range items {
	// 	fmt.Println(each)
	// 	wg.Add(1)
	// 	go sleep(each, n, tL, &wg)
	// }

	c := 0
	var m sync.Mutex

	for n, each := range items {
		fmt.Println(each)
		wg.Add(1)
		go sleep(each, n, tL, &wg, &c, &m)
	}

	// for n, each := range items {
	// 	fmt.Println(each)
	// 	wg.Add(1)
	// 	go sleep(each, n, tL, &wg, ch2)
	// 	tL = <-ch2
	// }
	wg.Wait()

}
