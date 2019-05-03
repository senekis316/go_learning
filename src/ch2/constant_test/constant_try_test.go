package constant_test

import "testing"

/**
  在其他语言中，连续常量的赋值一般如下：
  Monday = 1; Tuesday = 2; Wednesday = 3;
  但在go语言中，可以使用iota + 1这种方式来赋值
**/
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
