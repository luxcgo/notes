// 不把结构体的内容暴露到外面

package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type ma struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, and I am %d years old.\n", p.name, p.age)
}

func (p *ma) SayHello() {
	fmt.Println("我是69岁的老同志。")
}

func NewPerson(name string, age int) Person {
	if age == 69 {
		return &ma{"mabaoguo", 69}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Lucas", 22)
	p.SayHello()
	p2 := NewPerson("Tom", 69)
	p2.SayHello()
}
