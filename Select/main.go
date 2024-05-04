package main

import "fmt"

func addInt(ch chan int) {
	ch <- 111
}

func addFloats(ch chan float64) {
	ch <- 222.222
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan float64)
	go addInt(ch1)
	// fmt.Println(<-ch1)

	go addFloats(ch2)
	// fmt.Println(<-ch2)
	// select {
	// case data1 := <-ch1:
	// 	fmt.Println("data:", data1)
	// case data2 := <-ch2:
	// 	fmt.Println("data2:", data2)
	// }

	for i := range 2 {
		fmt.Println(i)
		select {
		case data1 := <-ch1:
			fmt.Println("data1:", data1)
		case data2 := <-ch2:
			fmt.Println("data2:", data2)
		}
	}
}
