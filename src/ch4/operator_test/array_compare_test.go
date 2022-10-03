package operator_test

import "testing"

const (
	Readable   = 1 << iota // 0001 1
	Writable               // 0010 2
	Executable             // 0100 4
	Delable                // 1000 8
)

type myInt int

func TestCompareArray(t *testing.T) {
	// 数组双等于号比较
	// 1. 数组的维度和值类型必须完全一致才能比较
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 3, 2, 4}
	// c := [...]myInt{1, 3, 2, 4}
	d := [...]int{1, 2, 3, 4, 5}

	t.Log(a == b) // true
	t.Log(a == c) // false
	t.Log(d)
	// t.Log(a == d)  长度不同,不同比较
}

func TestBitClear(t *testing.T) {
	// 按位与清零运算
	a := 7 // 0111
	t.Log(Readable)
	a = a &^ Readable
	t.Log(a)
	t.Log(a & Readable)
}
