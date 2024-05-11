package main

import (
	"fmt"
	"os"
)

type s struct {
	a, b int
}

func main() {
	a := 1
	fmt.Printf("a=%v\n", a)
	b := s{1, 2}
	fmt.Printf("struct: %v\n", b)
	fmt.Printf("struct with params: %+v\n", b)
	fmt.Printf("Struct #: %#v\n", b)
	fmt.Printf("Struct type: %T\n", b)
	fmt.Printf("boolean: %t\n", true)
	fmt.Printf("Int: %d\n", 999)
	fmt.Printf("binary of int %d: %b\n", 999, 999)
	fmt.Printf("Char of %d: %c\n", 123, 123)
	fmt.Printf("Hex of %d: %x\n", 1234, 1234)

	fmt.Printf("Float: %f\n", 1.23)
	data := "String-data"
	fmt.Printf("String: %s\n", data)
	fmt.Printf("String with quotes: \"%s\"\n", data)
	fmt.Printf("String Hex : %x\n", data)
	fmt.Printf("Pointer of var data: %p\n", &data)

	fmt.Printf("Width-Right: |%10d|%10d|\n", 1, 2)
	fmt.Printf("Width-Left: |%-10d|%-10d|\n", 1, 2)

	formattedString := fmt.Sprintf("Formatted data %s", data)
	fmt.Println(formattedString)

	fmt.Fprintf(os.Stderr, "io stdErr: %s", "errror")
}
