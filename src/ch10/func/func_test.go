package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//1.多返回值函数
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	//首次接受值，必须用赋值语法:=接收
	a, b := returnMultiValues()
	t.Log(a, b)

	//可以通过下滑线，来忽略一个返回值
	a, _ = returnMultiValues()
	t.Log(a)
}

//2.函数是一等公民
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn2(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

//3.可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log("Sum:", Sum(1, 2, 3, 4, 5))
}

//4.延迟运行defer函数
func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
}

func TestDeferError(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}()
	t.Log("Started")
	panic("Fatal Error") //defer仍会执行
}
