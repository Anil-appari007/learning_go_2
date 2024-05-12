package main

import (
	"fmt"
	"os/exec"
)

func checkErrAt(err error, at string) {
	if err != nil {
		fmt.Printf("Error: %s at %s\n", err.Error(), at)
	}
}

func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.CombinedOutput()
	checkErrAt(err, "date-CO")
	fmt.Println(string(dateOut))

	d2 := exec.Command("date")
	d2out, err := d2.Output()
	checkErrAt(err, "d2-O")
	fmt.Println(string(d2out))

	d3 := exec.Command("date", "+%Y")
	d3O, err := d3.Output()
	checkErrAt(err, "d3O")
	fmt.Println(string(d3O))

	d4 := exec.Command("dateeeee")
	d4O, err := d4.Output()
	checkErrAt(err, "d4-o")
	fmt.Println(string(d4O))
}
