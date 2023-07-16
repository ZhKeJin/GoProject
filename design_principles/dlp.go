package main

import (
	"fmt"
)

// 依赖倒置原则：描述的是，上层与下层之间的调用关系；
//			  上层调用下层不应该依赖于下层的具体的实现，而是依赖的于抽象（接口）；
//            核心思想是面向接口编程

type Cache interface {
	Set()
}

type Redis struct{}

func (r Redis) Set() {
	fmt.Println("redis")
}

type Abase struct{}

func (r Abase) Set() {
	fmt.Println("abase")
}

func main() {
	var cache Cache
	cache = Redis{}
	cache.Set()
	cache = Abase{}
	cache.Set()
}
