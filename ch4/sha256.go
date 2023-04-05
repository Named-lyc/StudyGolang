package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//Sum256函数对一个任意字节的slice类型的数据生成一个对应的信息摘要，大小为256bit，因此对应为[32]byte数组类型
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	//2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	//false
	//[32]uint8
}
