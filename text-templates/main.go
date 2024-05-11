package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func checkErr(err error, at string) {
	if err != nil {
		log.Printf("Error: %s at %s\n", err.Error(), at)
	}
}

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Hello {{.}}, Have a Great Day\n")
	checkErr(err, "t1-parse")
	err = t1.Execute(os.Stdout, "T1")
	checkErr(err, "t1-exec")

	t2 := template.Must(template.ParseFiles("t2.txt"))
	type User struct {
		Name string
	}
	u1 := User{"u1"}
	err = t2.Execute(os.Stdout, u1.Name)
	checkErr(err, "t2-exec")

	type Login struct {
		Name string
		Time time.Time
	}
	t3Login := Login{"t3", time.Now()}
	t3 := template.Must(template.ParseFiles("t3.txt"))
	err = t3.Execute(os.Stdout, t3Login)
	checkErr(err, "t3-exec")

	t4 := template.Must(template.ParseFiles("t4.txt"))
	users := []string{"dev1", "dev2", "dev3", "dev4"}
	err = t4.Execute(os.Stdout, users)

	t5 := template.Must(template.ParseFiles("t5.txt"))
	err = t5.Execute(os.Stdout, users)
	checkErr(err, "t5-exec")
}
