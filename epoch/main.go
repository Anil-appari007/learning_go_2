package main

import (
	"fmt"
	"time"
)

func main() {
	cTime := time.Now()
	p := fmt.Println
	p(cTime.Unix())
	p(cTime.UnixMilli())
	p(cTime.UnixMicro())
	p(cTime.UnixNano())

	p(time.Unix(1715454564, 0))
}
