package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (e *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, e.Position, e.AnnualIncome}
}

// functional approach
func NewEmployeeFactory(position string, annulIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annulIncome}
	}
}

func main() {
	// 方法一
	// 缺点：创建好Factory后无法修改
	// 优点：API角度传方法比传一个结构体更友好，传结构体的话调用方需要显示知道他有什么方法，不过这点可以通过interface改进。
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)
	developer := developerFactory("Adam")
	manager := managerFactory("Jane")
	fmt.Println(developer)
	fmt.Println(manager)

	// 方法二
	// 优点：创建好Factory后可以修改
	bossFactory := &EmployeeFactory{"CEO", 1000000}
	bossFactory.AnnualIncome++
	boss := bossFactory.Create("Steve")
	fmt.Println(boss)
}
