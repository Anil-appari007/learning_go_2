package main

import (
	"fmt"
	s "strings"
)

func main() {
	p := fmt.Println
	p("String Functions")
	p(s.ToUpper("hello"))

	p(s.ToLower("HELLOWORLD"))
	p(s.Title("testtitle"))
	p(s.Compare("test", "best"))
	p(s.Compare("test", "test"))
	p(s.Join([]string{"h", "e", "l", "l", "o", "j", "o", "i", "n"}, ""))
}
