package main

import (
	"fmt"
)

/*
Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
Printf 根据格式化参数生成格式化的字符串并写入标准输出。
*/
func main() {
	var stockcode = 123
	var enddate = "2023-6-15"
	var url = "Code=%dendDate=%s" //%d⇒整型數字 %s⇒字符串
	var url_string = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(url_string)
}
