package main

import (
	"flag"
	"fmt"

	math "github.com/WalterWJ/go-study/pkg/myMath/myMath"
)

func main() {
	num := flag.Int("n", 10, "number")
	flag.Parse()
	fmt.Println(*num)
	a = math.Add(1, 2)
}
