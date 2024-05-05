package main

import (
	"fmt"
	"time"
)

func workerThread(id int, tasks chan int, done chan int) {
	fmt.Println("Worker", id, "started")
	for each := range tasks {
		fmt.Println("Worker", id, "Received Task -", each)
		time.Sleep(2 * time.Second)
		fmt.Println("Worker", id, "completed Task ", each)

		done <- each
	}
}

func main() {
	count := 6
	tasks := make(chan int, count)
	done := make(chan int, count)

	for i := 1; i <= 2; i++ {
		go workerThread(i, tasks, done)
	}

	for j := 1; j <= count; j++ {
		tasks <- j
	}
	close(tasks)

	for d := 1; d <= count; d++ {
		<-done
		fmt.Println("Completed Tasks", d, "/", count)
	}
}
