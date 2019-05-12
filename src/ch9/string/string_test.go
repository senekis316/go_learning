package string_test

import (
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {

	//初始化默认值为""空字符串
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	//string是不可变的byte slice切片
	//s[1] = '3'

	s = "\xE4\xB8\xA5"
	t.Log(s)

	//len返回的是byte数
	s = "中"
	t.Log(len(s))

	//
	c := []rune(s)
	t.Log(len(c))

	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	//注意：在字符串转数字时，可能会出现无法转换的情况，所以转换结果是2个返回值。
	//如果没有err的返回，则代表数字转换成功
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
