package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	var c int = 8
	for c < 10 {
		c += 2
		fmt.Println(c)
	}

	var a = [5]int{2, 4, 6, 8, 10}
	for index, value := range a {
		fmt.Println(index, value)
	}
}
