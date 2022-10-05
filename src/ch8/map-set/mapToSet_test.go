package mapset_test

import "testing"

func TestMapToSetTry(t *testing.T) {
	//  使用map实现工厂模式
	m := map[string]func(num int) int{}

	m["getSelf"] = func(num int) int { return num }
	m["sqrtNum"] = func(num int) int { return num * num }

	// t.Logf("m[getSelf] type:%T", m["getSelf"])
}
