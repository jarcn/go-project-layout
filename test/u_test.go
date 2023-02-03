package test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"Hao", 44, 0, 8000}, {"Bob", 34, 10, 5000}, {"Alice", 23, 5, 9000}, {"Chenjia", 30, 8, 2000},
	{"Jack", 26, 0, 4000}, {"Tom", 48, 9, 7500}, {"Marry", 29, 0, 6000}, {"Mike", 32, 8, 4000},
}

func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i := range list {
		sum += fn(&list[i])
	}
	return sum
}

func TestMp(t *testing.T) {
	// 1）统计有多少员工大于 40 岁
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)

	//2）统计有多少员工薪水大于 6000
	high_pay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people: %d\n", high_pay)

	//3）列出有没有休假的员工
	no_vacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", no_vacation)

	//4）统计所有员工的薪资总和
	total_pay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})
	fmt.Printf("Total Salary: %d\n", total_pay)

	//5）统计 30 岁以下员工的薪资总和
	younger_pay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("Younger Pay: %d\n", younger_pay)
}
