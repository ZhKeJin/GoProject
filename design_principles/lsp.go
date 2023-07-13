package main

import "fmt"

// 里氏替换原则

type Action interface {
	Read()
}

type Book struct{}

func (b *Book) Read() {
	fmt.Println("Book..")
}

type Math struct {
	Book // 继承父类的方法
}

// 里氏替换，
func (m *Math) Read() {
	fmt.Println("Math....")
}

func user() {
	m := &Math{}
	m.Read()
}

func main() {
	user()
}
