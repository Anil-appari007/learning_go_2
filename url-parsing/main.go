package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	myUrl := "https://duckduckgo.com/?q=Go+Lang&t=brave&ia=web"
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(parsedUrl)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Protocol:", parsedUrl.Scheme)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Query", parsedUrl.Query())
}
