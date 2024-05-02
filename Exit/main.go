package main

import (
	"fmt"
	"os"
)

func customMessage(data string) {
	fmt.Println(data)
}

func main() {
	defer customMessage("Wont print")
	if 3 > 5 {
		fmt.Println("3 > 5")
		os.Exit(5)
	} else {
		customMessage("3<5")
		os.Exit(3)
	}

}
