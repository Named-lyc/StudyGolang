package main

import "fmt"

func main() {
	//数组是一段连续的内存块，在内存中顺序存储
	//数组长度固定
	//可以通过下标直接访问
	//如果没有初始化则会被默认初始化为元素类型的零值
	//数组的长度是数组的组成部分，因此[3]int和[4]int不是同一类型

	//数组的定义以及初始化
	var a [3]int             //array of 3 integers
	fmt.Println(a[0])        //print the a[0] element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	//打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	//只打印元素
	for _, v := range a {
		fmt.Printf("%d", v)
	}

	fmt.Println("数组的初始化：........")

	var b [3]int = [3]int{1, 2, 3}
	var c = [3]int{4, 5, 6}
	//...表示数组长度根据初始化的个数来确定
	q := [...]int{8, 9, 10, 11}
	fmt.Println(b[2])
	fmt.Println(c[2])
	fmt.Printf("%T\n", q)
	fmt.Println(len(q))

	fmt.Println("数组的比较。。。。。。")
	m := [2]int{1, 2}
	z := [...]int{1, 2}
	l := [...]int{1, 3}
	fmt.Println(m == z, z == l, m == l) //t,f,f
	//fmt.Println(m == q) mismatched types [3]int and [4]int
}
