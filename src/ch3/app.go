package main

import "fmt"

func main() {
	testImplicit()
	testPoint()
	testString()
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
	fmt.Printf("%T %T \n", a, aPtr)
}

func testString() {
	var s string
	// 字符串是数值类型,默认值为空字符,不是nil
	if len(s) > 1 {
		fmt.Println("ok")
	} else {

		fmt.Printf("no%s**", s)
	}
}
