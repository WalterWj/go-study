package main

import (
	"fmt"
)

type cb func(int) int

func main() {
	testCallBack(3, CallBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}

func testCallBack(x int, f cb) {
	x += 2
	f(x)
}

func CallBack(x int) int {
	fmt.Printf("Call back is %d \n", x)
	return x
}
