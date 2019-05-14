package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//1.组合方式实现继承
/*type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}*/

//2.匿名嵌套类方式实现继承

type Dog struct {
	Pet
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}
