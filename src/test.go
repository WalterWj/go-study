package main

import (
	"fmt"
	"regexp"

	mathClass "./myMath"
)

func main() {
	/* A simple code */
	fmt.Println("Hello, playground")
	str := "alter table aa add index IN_a(`a`)"
	reg := regexp.MustCompile("^alter\\s+table")
	data := reg.Find([]byte(str))
	fmt.Println(string(data))
	fmt.Println(mathClass.Add(1, 2))
	fmt.Println(mathClass.Sub(2, 1))
	const (
		a = 2 << iota
		b
	)

	fmt.Println(a, b)

	for m := 1; m < 10; m++ {
		/*    fmt.Printf("第%d次：\n",m) */
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}
}
