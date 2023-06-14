package main

//bool⇒布爾型
//int 整型
//int8 8位整型
//byte 字節型

import (
	sd "fmt"
)

type (
	byte uint8 //byte是uint8的別名
	rune int8
	テキスト string
)

func main() {
	a := 65
	b := string(a)
	sd.Println(b)
}
