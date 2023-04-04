package main

import "fmt"

// Celesius Fahrenheit有相同的底层类型float64，但是它们是不同的数据类型
type Celsius float64
type Fahrenheit float64 // 类型为Faharenheit的float64 但是不等价于float64

const (
	AbosluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC) !!type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) !! compile error : type mismatch
	fmt.Println(c == FToC(f))
}
