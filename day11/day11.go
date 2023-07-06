package main

import (
	"fmt"
)

type book struct {
	title  string
	auther string
	price  int
}

func main() {

	var book1 book
	var book2 book

	fmt.Println(book{"Golang言語", "johns", 100})
	fmt.Println(book{title: "Golang言語", auther: "johns", price: 100})

	book1.title = "Java"
	book1.auther = "不明"
	book1.price = 1500

	book2.title = "不明"
	book2.price = 5000

	fmt.Printf("Book1.title : %s\n", book1.title)
	fmt.Printf("Book2.price : %d\n", book2.price)
}
