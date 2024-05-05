package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./Logging/LogFile.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
	}
	log.SetOutput(file)

	log.Println("Hello - TestLog FIle 2")
}
