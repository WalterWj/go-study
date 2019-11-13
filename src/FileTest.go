package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()

	writeFile, err := os.OpenFile("test.txt", o.O_APPEND, 066)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(writeFile)
	writeFile.WriteString("test")
	writeFile.Close()
}
