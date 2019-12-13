package main

import (
	"flag"
	"fmt"

	myMath "github.com/WalterWJ/go-study/pkg/myMath"
)

func main() {
	num := flag.Int("n", 10, "number")
	flag.Parse()
	fmt.Println(*num)
	a = myMath.Add(1, 2)
}
