package main

import "fmt"

/*
	+	相加
	-	相减
	*	相乘
	/	相除
	%	求余
	++	自增
	--	自减
*/
func main() {
	/*
		+	相加
		-	相减
		*	相乘
		/	相除
		%	求余
		++	自增
		--	自减
	*/
	// var a int = 10
	// var b int = 20
	// var c int
	// c = a + b
	// fmt.Printf("a+bの値は %d\n", c)
	// c = b - a
	// fmt.Printf("b-aの値は %d\n", c)
	// c = a * b
	// fmt.Printf("a*bの値は %d\n", c)
	// c = b / a
	// fmt.Printf("b/aの値は %d\n", c)
	// c = b % a
	// fmt.Printf("baの値は %d\n", c)
	// a++
	// fmt.Printf("aの値は %d\n", a)
	// a--
	// fmt.Printf("aの値は %d\n", a)

	// /*

	// 	==	检查两个值是否相等，如果相等返回 True 否则返回 False。
	// 	!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。
	// 	>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。
	// 	<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。
	// 	>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
	// 	<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
	// */
	// var aa int = 20
	// var bb int = 30
	// if aa == bb {
	// 	fmt.Println("aaとbb一致")
	// }
	// if aa != bb {
	// 	fmt.Println("aa と　bb一致ではない")
	// }
	// if aa > bb {
	// 	fmt.Println("aa>bb")
	// }
	// if aa < bb {
	// 	fmt.Println("aa < bb")
	// }
	// if aa+10 >= bb {
	// 	fmt.Println("aa+10>=bb")
	// }
	// if aa <= bb-10 {
	// 	fmt.Println("aa <= bb-10")
	// }
	// /*
	// 	&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。
	// 	||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。
	// 	!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。

	// */
	// var aaa int = 20
	// var bbb int = 30
	// if aaa > bbb || aaa == 20 {
	// 	fmt.Println("||")
	// }
	// if aaa < bbb && aaa == 20 {
	// 	fmt.Println("&&")
	// }
	// if !(aaa > bbb) {
	// 	fmt.Println("!")
	// }

	/*
		&	返回变量存储地址
		*	指针变量。
	*/
	var aaaa int = 4
	// var bbbb int32
	// var cccc float32
	var pr *int //指針變量
	pr = &aaaa
	fmt.Printf("aaaaの値は%d\n", aaaa)
	fmt.Printf("pr保存了變量地址⇒%d\n", pr)
	fmt.Printf("*prは%d", *pr)

}
