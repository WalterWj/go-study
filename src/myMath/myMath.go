// myMath.go

package mathClass

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Max(x, y int) int {
	var result int

	if x > y {
		result = x
	} else {
		result = y
	}
	return result
}
