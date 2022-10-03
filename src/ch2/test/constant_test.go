package test_try

import "testing"

const (
	// iota 为go内置的特殊常量, 每次出现时值都会归0
	First = iota
	Second
	Third
	May
	July
)

func TestConstantTry(t *testing.T) {
	t.Log(First, Second, Third, May, July)
}
