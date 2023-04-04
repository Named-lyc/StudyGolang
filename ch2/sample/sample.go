package main

import "fmt"

func main() {
	var x, y int = 0, 0
	m := add(x, y)
	fmt.Println("x+y=", x+y)
	fmt.Println(m)
	var q *int
	//new初始化内存地址
	q = new(int)
	*q = 10
	fmt.Println(plus(q))
}

// 实参通过值的方式传递，因此函数的形参是实参的拷贝，
// 对形参进行修改不影响实参
func add(x, y int) int {
	x = 0
	y = 1
	return x + y
}

// 如果实参包含引用类型，那么拷贝的将是实参的地址，所以实参会由于函数的间接引用被修改
// 如 指针，slice，map，chan，function等
func plus(x *int) int {
	m := *x
	m += 1
	return m
}
