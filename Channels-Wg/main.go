package main

import (
	"fmt"
	"strconv"
	"time"
)

func worker(id int, toDo chan string, done chan string) {
	fmt.Printf("Worker started id: %v\n", id)
	c := 0
	for {
		select {
		case item := <-toDo:
			c++
			fmt.Printf("Worker %v received %s\n", id, item)
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker %v completed processing of %s\n", id, item)
			done <- item
		}
		fmt.Printf("Tasks completed by worker %v = %v\n", id, c)

	}
}

func main() {
	l := 30
	fmt.Printf("Tasks Count: %v\n", l)
	toDo := make(chan string, l)
	done := make(chan string, l)

	t := 3
	for each := range t {
		go worker(each, toDo, done)
	}

	for i := 1; i <= l; i++ {
		toDo <- strconv.Itoa(i)
	}
	close(toDo)
	cId := 0
	for range l {
		select {
		case doneTask := <-done:
			cId++
			fmt.Printf("%v/%v Task %s completed\n", cId, l, doneTask)
		}
	}
}
