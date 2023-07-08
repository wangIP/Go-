package main

/*
Go 语言切片(Slice)
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，
功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

Arrays配列
Goの Arrays(配列) は固定長の的な配列なので最初に宣言した配列のサイズを変えることはできません。

Slices
Goの Arrays(配列)は固定長の様な配列である一方、Goの Slices(スライス) は可変長の配列の様な動きをするのでより柔軟にデータ（要素）を格納することが可能です。
*/
import (
	"fmt"
)

// slice1 := make([]type,len)
func main() {
	//Array
	var arr1 [2]string
	arr1[0] = "Golang"
	arr1[1] = "Java"
	fmt.Println(arr1[0], arr1[1])

	var arr2 [2]string = [2]string{"Golang2", "Java2"}
	fmt.Println(arr2)

	//Slice
	var slice1 []string
	fmt.Println(slice1)

	slice2 := []string{"GolangS2", "JavaS2"}
	fmt.Println(slice2)

	arr3 := [...]string{"GolangARR", "JavaARR"}
	slice3 := arr3[0:1] //[Start:End]
	fmt.Println((slice3))

	slice3[0] = "Ruby"
	fmt.Println(slice3)

	//配列の長さが固定になるため、要素を追加する場合Sliceを使用する
	newSlice := append(slice3, "GoalngAdd")
	fmt.Println(newSlice)

	/*
		スライスは 長さ( length ) と 容量( capacity ) の両方を持っています。
		長さ(length) は、それに含まれる要素の数です。
		容量(capacity) は、スライスの最初の要素から数えて、元となる配列の要素数です。
	*/
	fmt.Println(slice3)      //Rubby
	fmt.Println(len(slice3)) //1⇒sliceの長さ
	fmt.Println(cap(slice3)) //2⇒sliceに代入された配列の長さ

	//make([]T,len,cap)
	//T⇒型,len⇒長さ,cap⇒容量
	a := make([]int, 5, 5)
	fmt.Println(a)

	//nil
	var number []int
	if number == nil {
		fmt.Println("slice is nil")
	}

	//copy
	b := []int{2, 3}
	bcopy := make([]int, 4, 4)
	copy(bcopy, b)
	fmt.Println(bcopy)

}
