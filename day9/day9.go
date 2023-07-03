package main

import (
	"fmt"
)

var a int = 10

/*配列初期化*/
var balance = [5]float32{100.0, 1.0, 2.0, 4.0, 5.0}

/*配列sizeが不明の場合*/
var balance1 = [...]int{1, 2, 3, 4, 5, 6}

/*indexで値をセット*/
var balance2 = [3]int{1: 20, 2: 30}

var s int = balance1[2]

func main() {
	/*
		Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
	*/
	/*var a int = 20*/
	var j int
	for j = 0; j < len(balance1); j++ {

		fmt.Printf("balance1[%d] = %d\n", j, balance1[j])
	}

}
