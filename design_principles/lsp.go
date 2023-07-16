package main

import "fmt"

// Action  里氏替换原则：描述的是子类方法不应该改变父类方法的定义；例如下面的 read()
type Action interface {
	Read()
}

type Book struct{}

func (b *Book) Read() {
	fmt.Println("Book..")
}

type Math struct {
	Book // *******继承父类的方法
}

// 里氏替换，
func (m *Math) Read() {
	fmt.Println("Math....")
}

func user() {
	m := &Math{}
	m.Read()
}

//func main() {
//	user()
//}
