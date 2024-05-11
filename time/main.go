package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p("Current-Time:", now)
	p("Year:", now.Year())
	p("Month:", now.Month())
	p("Date:", now.Day())

	pTime := time.Now()
	fTime := pTime.Format("2006-01-02 03:04:05PM -0700")
	p("Custom-Time:", fTime, pTime.Local())

	p(pTime.Format(time.RFC1123))
	p(pTime.Format(time.RFC1123Z))
	p(pTime.Format(time.RFC3339))
	p(pTime.Format(time.RFC3339Nano))

}
