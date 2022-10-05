package map_test

import "testing"

func TestMapDefine(t *testing.T) {
	// 3.声明并初始化
	m2 := map[string]int{"one": 1, "two": 2}

	// 4.make关键字 适合已知容量
	m3 := make(map[string]int, 10 /*cap*/)

	t.Log(m2, m3)
	// t.Log(cap(m2))  Map 不能使用cap()函数求容量
}

func TestAccessNotExistKey(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}

	// map中的key不存在返回0值,不能通过返回nil判断元素是否存在
	t.Log(m["three"]) // m[three] = 0

	// 判断key-value是否存在
	if v, ok := m["three"]; ok {
		t.Logf("m[three] = %d", v)
	} else {
		// 手动添加 m["three"] = 0
		m["three"] = 0
		t.Log(m)
	}

}
