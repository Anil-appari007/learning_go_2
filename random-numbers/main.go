package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	p := fmt.Println
	p(rand.IntN(100))
	p(rand.Float64())

	items := [6]string{"apple", "banana", "custard", "dragon", "pineapple", "strawberry"}
	p(items[rand.IntN(len(items))])
}
