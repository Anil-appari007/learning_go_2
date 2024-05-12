package main

import (
	"fmt"
	"os"
)

func main() {
	username := os.Getenv("USERNAME")
	fmt.Println(username)

	myOs := os.Getenv("OS")
	fmt.Println(myOs)

	allEnv := os.Environ()
	for _, each := range allEnv {
		fmt.Println(each)
	}

	dbPwd, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		fmt.Println("env var MYSQL_PASSWORD not found")
	} else {
		fmt.Println("dbPwd:", dbPwd)
	}

}
