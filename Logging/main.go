package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("./Logging/LogFile.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("Hello - TestLog FIle 2")

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog1 := slog.New(jsonHandler)
	slog1.Info("Hello From Logging")

	h2 := slog.NewTextHandler(os.Stdout, nil)
	slog2 := slog.New(h2)
	slog2.Info("Hello from slog2 NewTextHandler")

	logFile2, err := os.OpenFile("logFile-slog.txt", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer logFile2.Close()
	multiWriter := io.MultiWriter(os.Stdout, logFile2)
	h3 := slog.NewTextHandler(multiWriter, nil)
	slog3 := slog.New(h3)
	slog3.Info("Hello from slog3")
	slog3.Error("Errorrrr from s3")
	slog3.Info("CUstom attr test", "user", os.Getenv("USERNAME"))
}
