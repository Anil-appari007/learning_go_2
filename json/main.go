package main

import (
	"encoding/json"
	"fmt"
)

func checkErr(err error, at string) {
	if err != nil {
		fmt.Printf("Error: %s at %s", err.Error(), at)
	}
}

func main() {
	bol1, err := json.Marshal(true)
	checkErr(err, "bol1-marshall")
	fmt.Println(bol1)
	fmt.Println(string(bol1))

	int1, err := json.Marshal(123)
	checkErr(err, "int1-marshall")
	fmt.Println(int1)
	fmt.Println(string(int1))

	fl1, err := json.Marshal(4.56)
	checkErr(err, "fl1-marshall")
	fmt.Println(fl1)
	fmt.Println(string(fl1))
}
