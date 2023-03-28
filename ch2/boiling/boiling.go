package main

import "fmt"

// 包一级范围声明的可以在整个包下访问，及如果boiling包下还有个var A，则可以 var A= boilingF
const boilingF = 212.0

func main() {
	//f,c 在main包下声明，则其只能在main包下访问
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}
