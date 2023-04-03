package main

import "fmt"

func main() {
	/*new(type)会创建一个type类型的匿名变量，
	初始化type类型的零值，然后返回变量地址，
	返回类型为*type   */
	//new(type)*type
	p := new(int)   //p为*int类型，指向匿名的int变量
	fmt.Println(*p) //1
	*p = 2
	fmt.Println(*p) //2
	compare()

}

func compare() {
	//每次调用new函数都是返回一个新的变量的地址
	a := new(int)
	b := new(int)
	fmt.Println(a == b) //"false"
}
