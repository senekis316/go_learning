package type_test

import "testing"

//定义类型的别名
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//小类型向大类型转换，不丢失精度的情况下，go语言也不支持隐式类型转换
	b = int64(a)
	//go语言不支持任何的隐式类型转换，哪怕是同类型的别名都不行
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}
