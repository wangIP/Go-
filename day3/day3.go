package main

import (
	"fmt"
)

/*
iota

	每定義一個從0開始加1
	碰到const重置為0
*/
/*
6⇒2進数⇒ 0110
11⇒2進数⇒1011
&⇒兩個都是1的話就是1,有一個不是的話就是0⇒0010⇒2
|⇒只要有一個是1的話就是1⇒1111⇒15
^⇒兩位只有一個是1就是1⇒1101⇒13
&^⇒0100⇒4
*/
const (
	a = "A"
	b
	c
	d = iota
	e
)
const (
	f = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f + 1)
}
