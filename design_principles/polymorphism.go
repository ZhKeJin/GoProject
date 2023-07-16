package main

// DB 多态：go 多态是通过接口实现的; 提升代码的复用性，扩展性
type DB interface {
	Set()
	Read()
}

type Mysql struct{}

func (m Mysql) Set() {

}
func (m Mysql) Read() {

}

type Mongo struct{}

func (m Mongo) Set() {

}
func (m Mongo) Read() {

}

func GetDB(a int64) DB {
	switch a {
	case 1:
		return Mysql{}
	case 2:
		return Mongo{}
	}
	return Mysql{}
}

func Print(cache Cache) {
	cache.Set()
}

func case1() {
	db := GetDB(1)
	db.Read()
}

func case2() {
	m := Mysql{}
	Print(m)
}
