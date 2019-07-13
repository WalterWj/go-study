package main

import "fmt"

func main() {
	a, b := 100, 200
	fmt.Printf("Begin A and B is: %d and %d \n", a, b)

	swap(&a, &b)
	fmt.Printf("Change A and B is: %d and %d \n", a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
