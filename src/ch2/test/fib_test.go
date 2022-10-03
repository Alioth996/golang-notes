package test_try

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// 斐波那契数列 1 1 2 3 5 8 13
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		// tmp中间变量交换法
		// tmp := a
		// a = b
		// b = a + tmp

		// 多变量同时赋值并交换
		// a, b = b, a
		// b = a + b

		// 进一步简写
		a, b = b, a+b
	}

}
