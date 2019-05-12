package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

//1.创建结构体
type Employee struct {
	Id   string
	Name string
	Age  int
}

//2.通过结构体创建实例
func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //注意，这里返回的是指向实例的指针, 相当于e := &Employee{}
	e2.Id = "2"         //与其他主要编程语言的差异: 通过实例的指针访问成员不需要使用->
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e) //输出对象的类型
	t.Logf("e2 is %T", e2)
}

//3.定义方法(行为)
func (e Employee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) String2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e1 := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e1.Name))
	t.Log(e1.String1())

	e2 := &Employee{"10", "Jack", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e2.Name))
	t.Log(e2.String2())
}

/*func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}*/
