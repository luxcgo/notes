// 好处是结构体的有一些field值不变

package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// ..
	}
	return &Person{name, age, 2}
}

func main() {
	fmt.Println("dp")
	p := NewPerson("john", 33)
	log.Println("+v", p)
	p.EyeCount = 3
	p.Age++
	log.Println("+v", p)
}
