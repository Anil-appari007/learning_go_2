package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("Time-Out")

	time2 := time.AfterFunc(2*time.Second, func() { fmt.Println("Executing Function after 2 sec") })
	fmt.Println("Sleeping for 4 sec")
	time.Sleep(4 * time.Second)
	fmt.Println("Sleep done")
	time2.Stop()
}
