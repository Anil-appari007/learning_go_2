package main

import "fmt"

func pings(ch1 chan<- string, msg string) {
	ch1 <- msg
}

func pongs(ch1 <-chan string, ch2 chan<- string) {
	// msg := ""
	// msg <- ch1
	msg := <-ch1
	ch2 <- msg
}

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pings(ping, "hello")
	pongs(ping, pong)

	// fmt.Println(<-ping)
	fmt.Println(<-pong)

}
