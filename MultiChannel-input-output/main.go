package main

import (
	"log"
	"strconv"
	"time"
)

// Working
// func worker(id int, toDo chan string, done chan string) {
// 	log.Printf("Worker started id: %v\n", id)
// 	c := 0
// 	for {
// 		select {
// 		case item := <-toDo:
// 			c++
// 			log.Printf("Worker %v received %s\n", id, item)
// 			time.Sleep(2 * time.Second)
// 			log.Printf("Worker %v completed processing of %s\n", id, item)
// 			done <- item
// 		}
// 		log.Printf("Tasks completed by worker %v = %v\n", id, c)

// 	}
// }

/*
output of above
....
Worker 1 received
29/30 Task 30 completed
Worker 0 completed processing of 29
Tasks completed by worker 0 = 15
Worker 0 received
30/30 Task 29 completed

Issue: function is going through chan even after it is empty
*/

/////////////////////////////

func worker(id int, toDo chan string, done chan string) {
	log.Printf("Worker started id: %v\n", id)
	c := 0
	for item := range toDo {
		// select {
		// case item := <-toDo:
		c++
		log.Printf("Worker %v received task %s\n", id, item)
		time.Sleep(2 * time.Second)
		log.Printf("Worker %v completed processing of %s\n", id, item)
		done <- item
		// }
		log.Printf("Tasks completed by worker %v = %v\n", id, c)

	}
	log.Printf("")
}

/*

Tasks completed by worker 1 = 10
28/30 Task 29 completed
Worker 2 completed processing of 28
Tasks completed by worker 2 = 10
29/30 Task 28 completed
Worker 0 completed processing of 30
Tasks completed by worker 0 = 10
30/30 Task 30 completed
*/

// /////////////////////
func main() {

	/*
		1. Create a function which will take 2 channels
		2. Run function as go routine
		3. Give input to chan-1
		4. Send output from Function to Chan-2
	*/
	log.Println("Test")
	l := 30
	log.Printf("Tasks Count: %v\n", l)
	toDo := make(chan string, l)
	done := make(chan string, l)

	t := 7
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
			log.Printf("%v/%v Task %s completed\n", cId, l, doneTask)
		}
	}
}
