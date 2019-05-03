package string_test

import "testing"

func TestImplicit(t *testing.T) {
	//在其他主流语言中，String未初始化时是null或nil
	//而在go语言中，未初始化的String值为空字符串
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
