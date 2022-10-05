package map_test

import "testing"

func TestMapDefine(t *testing.T) {
	// 3.声明并初始化
	m2 := map[string]int{"one": 1, "two": 2}

	// 4.make关键字 适合已知容量
	m3 := make(map[string]int, 10 /*cap*/)

	t.Log(m2, m3)
	// t.Log(cap(m2))  Map 不能是cap()函数求容量
}
