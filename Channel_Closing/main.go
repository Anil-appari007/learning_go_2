package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := make(chan int, 5)
	taskDone := make(chan bool)

	go func() {
		for {
			task, more := <-tasks
			if more {
				fmt.Println("Received task", task)
			} else {
				fmt.Println("Received All Tasks")
				taskDone <- true
				return
			}
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Adding Task", i)
		tasks <- i
		fmt.Println("Added Task", i)
		time.Sleep(2 * time.Second)
	}
	close(tasks)
	fmt.Println("Tasks are added")
	<-taskDone
	_, ok := <-tasks
	fmt.Println("Received more tasks", ok)

	// IF data send to close channel
	// program will panic
	// tasks <- 333 // panic: send on closed channel
}
