package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		// if n == 1 {
		// 	t.Log(n)
		// 	break
		// }

		// if n == 2 {
		// 	t.Log(n)
		// 	continue
		// }
		t.Log(n)
		n++
	}
}
