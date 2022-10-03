package main

import "fmt"

func main() {
	testImplicit()
	testPoint()
}

type MyInt int64

// 类型转换
func testImplicit() {
	var a int32 = 1
	var b int64
	var c MyInt
	b = int64(a)
	c = MyInt(b)
	fmt.Println(a, b, c)
}

// 指针
func testPoint() {
	a := 1
	aPtr := &a
	// aPtr += 1 指针不能运算
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
}
