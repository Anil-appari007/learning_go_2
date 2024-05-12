package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	demo := "Password"
	h1 := sha256.New()
	h1.Write([]byte(demo))
	fmt.Printf("h1: %x\n", h1.Sum(nil))

	h2 := sha256.Sum256([]byte(demo))
	fmt.Printf("h2: %x", h2)
}
