package test_try

import "testing"

func TestFirstTry(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a // 交换两个变量的值
	t.Log("my fisrt test try")
	t.Logf("a:%d,b:%d", a, b)
}
