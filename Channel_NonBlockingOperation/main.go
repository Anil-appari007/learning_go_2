package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	// time.Sleep(3 * time.Second)

	select {
	case ch1 <- 1111111:
		fmt.Println("data sent to channel")
	default:
		fmt.Println("Channel is not ready to receive")
	}

	select {
	case msg := <-ch1:
		fmt.Println("Data received from Channel, ", msg)
	default:
		fmt.Println("Channel is not ready to send data")
	}

	ch2 := make(chan int, 1)
	select {
	case ch2 <- 22222:
		fmt.Println("data Sent to ch2")
	default:
		fmt.Println("Chanel ch2 is not ready to receive")
	}

	select {
	case msg2 := <-ch2:
		fmt.Printf("message received from Channel %v, data is %v.\n", ch2, msg2)
	default:
		fmt.Printf("Channel is not ready to send data.\n")
	}
}
