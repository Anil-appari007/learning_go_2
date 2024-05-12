package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkErrAt(err error, at string) {
	if err != nil {
		fmt.Printf("Error: %s at %s\n", err.Error(), at)
	}
}
func main() {
	res, err := http.Get("http://ipv4.icanhazip.com")
	checkErrAt(err, "Http-get")
	defer res.Body.Close()
	fmt.Println(res)
	fmt.Println(res.Body)

	data, err := io.ReadAll(res.Body)
	checkErrAt(err, "io.readAll")
	fmt.Println(string(data))
}
