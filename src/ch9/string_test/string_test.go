package string_test

import (
	"testing"
)

func TestStringTry(t *testing.T) {
	s := "hello"

	// srting 不可变
	// s[1] = "ll"
	//s = "\xE4\xB8\xA5" // 可以存储二进制数据
	//t.Log(s, len(s))   // s = 严,len=3,3段编码

	s = "中"
	t.Log(len(s)) // len = 3,len获取的是底层 byte 的长度 utf8:e4b8ad

	// 转换
	c := []rune(s)
	t.Log(len(c)) // len = 1,因为s只有一个中 字
	// t.Log("c :", c)
	// 取16进制编码值
	t.Logf("中 unicode: %x", c[0])
	t.Logf("中 utf8:%x", s)
}

func TestStringToRune(t *testing.T) {
	s := "李时珍的皮"

	for _, v := range s {
		// 输出:
		// 李 674e
		// 时 65f6
		// 珍 73cd
		// 的 7684
		// 皮 76ae
		t.Logf("%c:%x", v, v)
	}
}
