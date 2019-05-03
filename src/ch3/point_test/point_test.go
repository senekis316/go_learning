package point_test

import "testing"

func TestImplicit(t *testing.T) {
	a := 1
	aPtr := &a
	//go语言不支持对指针值的运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
