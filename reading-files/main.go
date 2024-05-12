package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func checkErrAt(err error, at string) {
	if err != nil {
		log.Printf("Error: %s at %s\n", err.Error(), at)
	}
}

func createLogger(file string) *os.File {
	logFile, err := os.OpenFile(file, os.O_CREATE, 0666)
	checkErrAt(err, "createLogger-openfile")
	return logFile
}

func customLogger(data string, logFile *os.File) {
	log.Println(data)
	logFile.Write([]byte(data))
}

func main() {
	f1, err := os.OpenFile("f1.txt", os.O_CREATE, 666)
	checkErrAt(err, "f1-openfile")
	defer f1.Close()
	fmt.Println(f1.Stat())
	f1.Write([]byte("Sample Data"))
	// f1.WriteAt([]byte("data2"), 3)

	// Aim create func which will write logs to file and also prints to stdout
	f2 := createLogger("f2.txt")
	customLogger("hello from log2", f2)

	f3, err := os.OpenFile("f3.txt", os.O_CREATE, 0600)
	checkErrAt(err, "f3-openfile")
	defer f3.Close()
	logFile := io.MultiWriter(f3, os.Stdout)
	log.SetOutput(logFile)
	log.Println("MultiWriter-f3")
}
