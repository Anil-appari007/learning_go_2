package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkErrAt(err error, at string) {
	if err != nil {
		fmt.Printf("Error: %s at %s\n", err.Error(), at)
	}
}

func main() {
	cd, err := os.Getwd()
	checkErrAt(err, "cd")
	fmt.Println("Current Dir:", cd)

	items, err := os.ReadDir(cd)
	checkErrAt(err, "io")
	// fmt.Println(items)
	for i, item := range items {
		info := item.Type()
		// info, err := item.Info()
		// checkErrAt(err, "info")
		fmt.Printf("%v. %s - %s\n", i, item.Name(), info)
	}

	multiDir := filepath.Join("test1", "test2")
	fmt.Println(multiDir)

	err = os.Mkdir("d1", 400)
	checkErrAt(err, "mkdir d1")
	os.Remove("d1")
}
