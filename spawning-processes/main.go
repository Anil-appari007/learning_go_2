package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	// m1Params := "-uroot -p -e SHOW DATABASES"

	m1 := exec.Command("mysql", "-uroot", "-p", "-e", "SHOW DATABASES LIKE 'db%'")
	m1O, err := m1.CombinedOutput()
	fmt.Println(string(m1O))

	m2Input, err := os.ReadFile("D:\\CODE_SAMPLEs\\Github\\learning_go_2\\spawning-processes\\data.sql")
	checkErrAt(err, "m2Input")
	fmt.Println(string(m2Input))

	m2 := exec.Command("mysql", "-uroot", "-p", "-f")
	// m2.Stdin = string(m2Input)
	m2.Stdin = strings.NewReader(string(m2Input))
	m2O, err := m2.CombinedOutput()
	checkErrAt(err, "m2O")
	fmt.Println(string(m2O))

}
