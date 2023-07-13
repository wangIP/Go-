package main

import (
	"fmt"
)

// map
func main() {
	/*
		m := make(map[string]int)  key型⇒String,Value型int
		m := make(map[string]int,10) key型⇒String,Value型int,容量10
		m := map[string]int{
			"apple":100,
			"Orange":200,
			"banana":150
		}
	*/
	//宣言
	var m map[string]string
	//初期化
	m = make(map[string]string)
	m["name"] = "john"
	m["Contury"] = "TW"
	m["City"] = "TYP"

	for mm := range m {
		fmt.Println(mm, " is", m[mm])
	}
	delete(m, "City")
	for mm := range m {
		fmt.Println(mm, " is", m[mm])
	}

	//宣言と初期化一緒
	ms := make(map[string]int)
	//ms = make(map[string]string)
	ms["price"] = 200
	fmt.Println(ms["price"])
	//////////////////////////////
}
