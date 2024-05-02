package main

import (
	"fmt"
	"sync"
)

func sayHelloNTimes(s string) {
	for each := range 10 {
		fmt.Println(each, "HEllo", s)
	}
}

func goSayHello(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for each := range 10 {
		fmt.Println(each, "Hello from Routine", s)
	}
}

func main() {
	sayHelloNTimes("User")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		sayHelloNTimes("Go-Routine")
	}()

	wg.Add(1)
	go goSayHello("goUser", &wg)
	wg.Wait()

	// time.Sleep(10 * time.Second)
}
