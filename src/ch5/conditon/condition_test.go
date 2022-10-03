package conditon_test

import "testing"

func TestMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}
}

func TestSwitchCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		//  没有break语句
		switch i {
		// case 可以多项选一,
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}

	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := -1; i < 5; i++ {
		// 注意: 此处 switch 后面没有条件表达式,则case语句需要进行条件判断
		switch {
		case i%2 == 0:
			t.Log("even")

		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unkown")
		}

	}
}
