package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hello, playground")
	str :="alter table aa add index IN_a(`a`)"
	reg :=regexp.MustCompile("^alter\\s+table")
	data := reg.Find([]byte(str))
	fmt.Println(string(data))
	fmt.Println("Hello, World!")
}
