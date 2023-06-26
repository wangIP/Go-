package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 20
	var y int = 30
	fmt.Println(funcation_max(x, y))

	fmt.Println(funcation_nameStr("john", 123456))

	z := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(z(9))
}

// Go 函数可以返回多个值
func funcation_max(x, y int) int {
	var result int
	if x > y {
		result = x
	} else {
		result = y
	}
	return result
}

// Go 函数可以返回多个不一樣型態的值
func funcation_nameStr(name1 string, phone int) (string, int) {
	return name1, phone
}
