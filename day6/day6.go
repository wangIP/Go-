package main

import (
	"fmt"
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

/*
たとえば十進数の5を二進数であらわすと101になります。
そこで左にずらしてみます。
101 << 1 → 1010(二進数)で十進数だと10です。
101 << 3 → 101000(二進数)で十進数だと40です。

今度は右にずらしてみます。
101 >> 1 → 10(二進数)で十進数だと2です。
101 >> 2 → 1(二進数)で十進数でも1です。
*/
func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
