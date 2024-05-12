package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	myPwd := "goLang@!#$()password"
	bE := b64.StdEncoding.EncodeToString([]byte(myPwd))
	fmt.Printf("Encrypted form of %s: %s\n", myPwd, bE)
	bD, _ := b64.StdEncoding.DecodeString(bE)
	fmt.Printf("Decrypted: %s\n", bD)

	uE := b64.URLEncoding.EncodeToString([]byte(myPwd))
	fmt.Printf("URL Excryption: %s\n", uE)
	uD, _ := b64.URLEncoding.DecodeString(uE)
	fmt.Printf("URL Decryption: %s", uD)

}
